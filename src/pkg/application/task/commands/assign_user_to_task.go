package commands

// AssignUserToTaskInput ...
type AssignUserToTaskInput struct {
	AssigneeID string `json:"assignee_id"`
	TaskID     string `json:"task_id"`
}

// AssignUserToTaskOutput ...
type AssignUserToTaskOutput struct {
	TaskDescriptor
}
