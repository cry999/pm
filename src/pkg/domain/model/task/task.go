package task

import (
	"context"
	"time"

	"github.com/cry999/pm-projects/pkg/domain/errors/common"
	"github.com/cry999/pm-projects/pkg/domain/event"
	taskEvent "github.com/cry999/pm-projects/pkg/domain/event/task"
)

// Task ...
type Task struct {
	ID          ID
	Name        string
	Description string
	OwnerID     UserID
	ProjectID   ProjectID
	AssigneeID  UserID
	Status      Status
	Deadline    *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// New creates a new Task instance
func New(id ID, name, description string, ownerID UserID) (_ *Task, err error) {
	t := new(Task)
	if err = t.setID(id); err != nil {
		return
	}
	if err = t.setName(name); err != nil {
		return
	}
	if err = t.setDescription(description); err != nil {
		return
	}
	if err = t.setOwnerID(ownerID); err != nil {
		return
	}
	t.AssigneeID = UserIDZero
	t.Status = StatusTodo

	return t, nil
}

// PlanForProject ...
func PlanForProject(id ID, name, description string, ownerID UserID, projectID ProjectID) (_ *Task, err error) {
	t, err := New(id, name, description, ownerID)
	if err != nil {
		return
	}
	if err = t.SetProjectID(projectID); err != nil {
		return
	}
	event.Get().Publish(context.Background(), taskEvent.Planned{
		ID:          t.ID.String(),
		Name:        t.Name,
		Description: t.Description,
		OwnerID:     t.OwnerID.String(),
		ProjectID:   t.ProjectID.String(),
		AssigneeID:  t.AssigneeID.String(),
		Status:      string(t.Status),
		Deadline:    t.Deadline,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	})
	return t, nil
}

// setID ...
func (t *Task) setID(id ID) error {
	if id.Equals(IDZero) {
		return common.InvalidArgumentError("id", "zero")
	}
	if !t.ID.Equals(IDZero) {
		return common.IllegalOperationError("user.id is already set")
	}
	t.ID = id
	return nil
}

// setName ...
func (t *Task) setName(name string) error {
	if name == "" {
		return common.InvalidArgumentError("name", "empty")
	}
	t.Name = name
	return nil
}

// setDescription ...
func (t *Task) setDescription(description string) error {
	if description == "" {
		return common.InvalidArgumentError("description", "empty")
	}
	t.Description = description
	return nil
}

// setOwnerID ...
func (t *Task) setOwnerID(ownerID UserID) error {
	if ownerID.Equals(UserIDZero) {
		return common.InvalidArgumentError("owner_id", "zero")
	}
	if !t.OwnerID.Equals(UserIDZero) {
		return common.IllegalOperationError("owner_id is already set")
	}
	t.OwnerID = ownerID
	return nil
}

// SetProjectID ...
func (t *Task) SetProjectID(projectID ProjectID) error {
	if !ProjectIDZero.Equals(t.ProjectID) {
		return common.IllegalOperationError("this task already belongs to the project")
	}
	t.ProjectID = projectID
	return nil
}

// Assign ...
func (t *Task) Assign(assigneeID UserID) error {
	if assigneeID.Equals(UserIDZero) {
		return common.InvalidArgumentError("assignee_id", "zero")
	}
	if assigneeID.Equals(t.AssigneeID) {
		// same user is already assigned to this task
		return nil
	}
	if !t.AssigneeID.Equals(UserIDZero) {
		// different user is already assigned to this task
		t.Unassign()
	}
	t.AssigneeID = assigneeID
	return nil
}

// Unassign ...
func (t *Task) Unassign() {
	t.AssigneeID = UserIDZero
}

// SetDeadline ...
func (t *Task) SetDeadline(deadline time.Time) error {
	t.Deadline = &deadline
	return nil
}

// UnsetDeadline ...
func (t *Task) UnsetDeadline() {
	t.Deadline = nil
}

// Progress ...
func (t *Task) Progress() error {
	if t.Status != StatusTodo {
		return t.illegalStatusError(StatusTodo, StatusWip)
	}
	t.Status = StatusWip
	return nil
}

// Done ...
func (t *Task) Done() error {
	if t.Status != StatusWip {
		return t.illegalStatusError(StatusWip, StatusDone)
	}
	t.Status = StatusDone
	return nil
}

// Redo ...
func (t *Task) Redo() error {
	t.Status = StatusTodo
	return nil
}

// PutOnHold ...
func (t *Task) PutOnHold() error {
	t.Status = StatusPending
	return nil
}

// Cancel ...
func (t *Task) Cancel() error {
	t.Status = StatusCancel
	return nil
}

func (t *Task) illegalStatusError(allowed, next Status) error {
	return common.IllegalOperationError(
		"only '%s' status task can be '%s'. This task is '%s'",
		allowed, next, t.Status,
	)
}

// CanAccess ...
func (t *Task) CanAccess(userID UserID) bool {
	return userID.Equals(t.OwnerID) || userID.Equals(t.AssigneeID)
}
