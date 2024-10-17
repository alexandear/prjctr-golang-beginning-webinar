package http

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/iancoleman/strcase"
)

var (
	fieldNestPattern = regexp.MustCompile(`\[\d+\]$`)
	ErrWrongID       = errors.New("wrong id")
)

type ValidationError struct {
	Field string `json:"field"`
	Error string
}

type ErrorResponse struct {
	Code    int                 `json:"code"`
	Errors  map[string][]string `json:"errors,omitempty"`
	Message string              `json:"message"`
}

func toField(field validator.FieldError) string {
	var parts []string
	for _, part := range strings.Split(field.StructNamespace(), ".")[1:] {
		part = strcase.ToSnake(part)
		func() { // separate fieldName[N] to fieldName and N
			nested := fieldNestPattern.FindString(part)
			if nested != "" {
				part = strings.ReplaceAll(part, nested, "")
				defer func() { parts = append(parts, nested[1:len(nested)-1]) }()
			}
			parts = append(parts, strcase.ToSnake(part))
		}()
	}

	return "[" + strings.Join(parts, "][") + "]"
}

func NewErrorResponse(c *gin.Context, err error) {
	var ve validator.ValidationErrors

	resp := &ErrorResponse{
		Code:    0,
		Message: "Unprocessed error",
	}

	if errors.As(err, &ve) {
		resp.Code = http.StatusBadRequest
		resp.Errors = make(map[string][]string)
		resp.Message = "Got errors while validation"

		for _, fieldErr := range ve {
			resp.Errors[toField(fieldErr)] = append(resp.Errors[toField(fieldErr)], fmt.Sprintf("Field validation failed on the '%s' tag", fieldErr.ActualTag()))
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	if uuid.IsInvalidLengthError(err) || errors.Is(err, ErrWrongID) || strings.Contains(err.Error(), "is required parameter") {
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = http.StatusInternalServerError
	c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
}
