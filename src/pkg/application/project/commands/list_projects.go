package commands

// ListProjectsInput ...
type ListProjectsInput struct {
	UserID string `json:"user_id"`
}

// ListProjectsOutput ...
type ListProjectsOutput struct {
	Results []ProjectDescriptor `json:"results"`
}
