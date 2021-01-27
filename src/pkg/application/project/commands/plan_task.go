package commands

import "github.com/cry999/pm-projects/pkg/application/task/commands"

// TODO: task domain に入れるべきか要検討

// PlanTaskInput ...
type PlanTaskInput struct {
	OwnerID     string `json:"owner_id"`
	ProjectID   string `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// PlanTaskOutput ...
type PlanTaskOutput struct {
	commands.TaskDescriptor
}
