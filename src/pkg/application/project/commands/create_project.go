package commands

// CreateProjectInput ...
type CreateProjectInput struct {
	OwnerID       string `json:"owner_id"`
	Name          string `json:"name"`
	ElevatorPitch string `json:"elevator_pitch"`
}

// CreateProjectOutput ...
type CreateProjectOutput struct {
	ProjectDescriptor
}
