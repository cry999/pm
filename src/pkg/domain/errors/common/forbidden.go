package common

import "fmt"

// Forbidden ...
type Forbidden struct {
	actor string
	role  string
}

// ForbiddenError ...
func ForbiddenError(actor, role string) *Forbidden {
	return &Forbidden{
		actor: actor,
		role:  role,
	}
}

// Error ...
func (e *Forbidden) Error() string {
	return fmt.Sprintf("'%s' has no role '%s'", e.actor, e.role)
}
