package extend

import "fmt"

type FormattedJsonError struct {
	Code    int
	Message string
	Err     error // вкладена помилка
}

func (e *FormattedJsonError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf(`{"%d": "%s: %v"}`, e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf(`{"%d": "%s"}`, e.Code, e.Message)
}

func NewFormattedError(code int, message string, err error) *FormattedJsonError {
	return &FormattedJsonError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
