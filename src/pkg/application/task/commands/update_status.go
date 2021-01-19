package commands

var (
	// Actions ...
	Actions = []string{
		"progress",
		"done",
		"redo",
		"cancel",
		"hold",
	}
)

// UpdateStatusInput ...
type UpdateStatusInput struct {
	ActorID string `json:"actor_id"`
	TaskID  string `json:"task_id"`
	Action  string `json:"action"`
}

// UpdateStatusOutput ...
type UpdateStatusOutput struct {
	TaskDescriptor
}
