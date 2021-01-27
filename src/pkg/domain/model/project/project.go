package project

import (
	"time"

	"github.com/cry999/pm-projects/pkg/domain/errors/common"
	"github.com/cry999/pm-projects/pkg/domain/model/task"
)

// Project ...
type Project struct {
	ID            ID
	OwnerID       UserID
	Name          string
	ElevatorPitch string
	PlannedTasks  []PlannedTaskID
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// New creates a new Project instance
func New(id ID, ownerID UserID, name, elevatorPitch string) (_ *Project, err error) {
	p := new(Project)
	if err = p.setID(id); err != nil {
		return
	}
	if err = p.setOwnerID(ownerID); err != nil {
		return
	}
	if err = p.setName(name); err != nil {
		return
	}
	if err = p.setElevatorPitch(elevatorPitch); err != nil {
		return
	}
	return p, nil
}

// setID ...
func (p *Project) setID(id ID) error {
	if IDZero.Equals(id) {
		return common.InvalidArgumentError("id", "zero")
	}
	if !IDZero.Equals(p.ID) {
		return common.IllegalOperationError("id is already set")
	}
	p.ID = id
	return nil
}

// setOwnerID ...
func (p *Project) setOwnerID(ownerID UserID) error {
	if UserIDZero.Equals(ownerID) {
		return common.InvalidArgumentError("owner_id", "zero")
	}
	p.OwnerID = ownerID
	return nil
}

// setName ...
func (p *Project) setName(name string) error {
	if name == "" {
		return common.InvalidArgumentError("name", "empty")
	}
	p.Name = name
	return nil
}

// setElevatorPitch ...
func (p *Project) setElevatorPitch(elevatorPitch string) error {
	if elevatorPitch == "" {
		return common.InvalidArgumentError("elevator_pitch", "empty")
	}
	p.ElevatorPitch = elevatorPitch
	return nil
}

// PlanTask ...
func (p *Project) PlanTask(taskID task.ID, name, description string, ownerID task.UserID) (*task.Task, error) {
	taskProjectID, err := task.NewProjectID(p.ID.String())
	if err != nil {
		return nil, err
	}
	t, err := task.PlanForProject(taskID, name, description, ownerID, taskProjectID)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// PlannedTask ...
func (p *Project) PlannedTask(plannedTaskID PlannedTaskID) error {
	p.PlannedTasks = append(p.PlannedTasks, plannedTaskID)
	return nil
}
