package pubsub

import (
	"context"
	"io"
)

// Subscriber ...
type Subscriber func(ctx context.Context, payload io.Reader) error

// Middleware ...
type Middleware func(next Subscriber) Subscriber
