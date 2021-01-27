package commands

// RetrieveProjectInput ...
type RetrieveProjectInput struct {
	OwnerID   string `json:"owner_id"`
	ProjectID string `json:"project_id"`
}

// RetrieveProjectOutput ...
type RetrieveProjectOutput struct {
	ProjectDescriptor
}
