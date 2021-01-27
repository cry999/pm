package commands

import (
	"github.com/cry999/pm-projects/pkg/application/task/commands"
	"github.com/cry999/pm-projects/pkg/domain/model/project"
	"github.com/cry999/pm-projects/pkg/domain/model/task"
)

// ProjectDescriptor ...
type ProjectDescriptor struct {
	ID            string                    `json:"id"`
	OwnerID       string                    `json:"owner_id"`
	Name          string                    `json:"name"`
	ElevatorPitch string                    `json:"elevator_pitch"`
	Tasks         []commands.TaskDescriptor `json:"tasks"`
}

// NewProjectDescriptor ...
func NewProjectDescriptor(p *project.Project, tasks []*task.Task) ProjectDescriptor {
	desc := ProjectDescriptor{
		ID:            p.ID.String(),
		OwnerID:       p.OwnerID.String(),
		Name:          p.Name,
		ElevatorPitch: p.ElevatorPitch,
		Tasks:         []commands.TaskDescriptor{},
	}
	for _, t := range tasks {
		desc.Tasks = append(desc.Tasks, commands.NewTaskDescriptor(t))
	}
	return desc
}
