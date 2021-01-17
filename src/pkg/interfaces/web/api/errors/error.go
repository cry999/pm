package errors

import "fmt"

// HTTPError ...
type HTTPError struct {
	Code    int
	Message string
}

// HTTPErrorf ...
func HTTPErrorf(code int, f string, a ...interface{}) *HTTPError {
	return &HTTPError{
		Code:    code,
		Message: fmt.Sprintf(f, a...),
	}
}

// Error ...
func (e *HTTPError) Error() string { return e.Message }
