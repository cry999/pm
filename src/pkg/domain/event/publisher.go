package event

import (
	"context"
)

// Publisher ...
type Publisher interface {
	Publish(ctx context.Context, event Event) error
}

var (
	registered Publisher
)

// Register ...
func Register(publisher Publisher) {
	if registered != nil {
		panic("publisher is already registered")
	}
	registered = publisher
}

// Get ...
func Get() Publisher {
	if registered == nil {
		panic("publisher is not registered")
	}
	return registered
}
