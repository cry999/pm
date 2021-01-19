package task

// Status of Tasks
type Status string

// Status enum
const (
	StatusUnknown Status = "UNKNOWN"
	StatusTodo    Status = "TODO"
	StatusWip     Status = "WIP"
	StatusDone    Status = "DONE"
	StatusPending Status = "PENDING"
	StatusCancel  Status = "CANCEL"
)
