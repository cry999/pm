package event

// Event ...
type Event interface {
	// Type of the Event MUST be unique in the global
	Type() string
}
