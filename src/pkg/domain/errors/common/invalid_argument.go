package common

import "fmt"

// InvalidArgument ...
type InvalidArgument struct {
	field   string
	message string
}

// InvalidArgumentError ...
func InvalidArgumentError(field, format string, a ...interface{}) *InvalidArgument {
	return &InvalidArgument{
		field:   field,
		message: fmt.Sprintf(format, a...),
	}
}

// Error ...
func (e *InvalidArgument) Error() string { return fmt.Sprintf("%s: %s", e.field, e.message) }
