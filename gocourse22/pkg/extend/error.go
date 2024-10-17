package extend

import "fmt"

type FormattedJSONError struct {
	Code    int
	Message string
	Err     error // вкладена помилка
}

func (e *FormattedJSONError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf(`{"%d": "%s: %v"}`, e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf(`{"%d": "%s"}`, e.Code, e.Message)
}

func NewFormattedError(code int, message string, err error) *FormattedJSONError {
	return &FormattedJSONError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
