package commands

// RetrieveTaskInput ...
type RetrieveTaskInput struct {
	UserID string `json:"user_id"`
	TaskID string `json:"task_id"`
}

// RetrieveTaskOutput ...
type RetrieveTaskOutput struct {
	TaskDescriptor
}
