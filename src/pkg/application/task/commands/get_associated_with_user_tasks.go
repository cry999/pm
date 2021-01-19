package commands

// GetAssociatedWithUserTasksInput ...
type GetAssociatedWithUserTasksInput struct {
	UserID string `json:"user_id"`
}

// GetAssociatedWithUserTasksOutput ...
type GetAssociatedWithUserTasksOutput struct {
	Results []TaskDescriptor `json:"results"`
}
