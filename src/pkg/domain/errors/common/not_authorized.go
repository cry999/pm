package common

// NotAuthorized ...
type NotAuthorized struct {
}

// NotAuthorizedError ...
func NotAuthorizedError() *NotAuthorized {
	return &NotAuthorized{}
}

// Error ...
func (e *NotAuthorized) Error() string {
	return "not authorized"
}
