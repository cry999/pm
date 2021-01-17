package common

import "fmt"

// NotFound ...
type NotFound struct {
	Entity string
	Key    string
}

// NotFoundError ...
func NotFoundError(entity, key string) *NotFound {
	return &NotFound{
		Entity: entity,
		Key:    key,
	}
}

// Error ...
func (e *NotFound) Error() string {
	return fmt.Sprintf("not found: %s(key=%s)", e.Entity, e.Key)
}
