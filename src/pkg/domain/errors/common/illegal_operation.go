package common

import "fmt"

// IllegalOperation ...
type IllegalOperation struct {
	message string
}

// IllegalOperationError creates a new IllegalOperation instance
func IllegalOperationError(f string, a ...interface{}) *IllegalOperation {
	return &IllegalOperation{
		message: fmt.Sprintf(f, a...),
	}
}

// Error ...
func (e *IllegalOperation) Error() string { return fmt.Sprintf("illegal operation: %s", e.message) }
