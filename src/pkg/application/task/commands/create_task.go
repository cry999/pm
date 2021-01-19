package commands

// CreateTaskInput ...
type CreateTaskInput struct {
	OwnerID     string `json:"owner_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateTaskOutput ...
type CreateTaskOutput struct {
	TaskDescriptor
}
