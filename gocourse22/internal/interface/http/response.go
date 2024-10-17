package http

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int             `json:"-"`
	Data       json.RawMessage `json:"data,omitempty"`
}

func NewResponse(c *gin.Context, data any, opts ...ResponseOpt) {
	jsn, err := json.Marshal(data)
	if err != nil {
		NewErrorResponse(c, err)
		return
	}

	r := &Response{
		StatusCode: http.StatusOK,
		Data:       jsn,
	}

	for _, opt := range opts {
		opt(r)
	}

	c.JSON(r.StatusCode, r)
}

type ResponseOpt func(r *Response)

func WithStatusCode(statusCode int) ResponseOpt {
	return func(r *Response) {
		r.StatusCode = statusCode
	}
}
