package task

import (
	"time"

	"github.com/cry999/pm-projects/pkg/domain/errors/common"
)

// Task ...
type Task struct {
	ID          ID
	Name        string
	Description string
	OwnerID     UserID
	AssigneeID  UserID
	Status      Status
	Deadline    *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewTask creates a new Task instance
func NewTask(id ID, name, description string, ownerID UserID) (_ *Task, err error) {
	t := new(Task)
	if err = t.SetID(id); err != nil {
		return
	}
	if err = t.SetName(name); err != nil {
		return
	}
	if err = t.SetDescription(description); err != nil {
		return
	}
	if err = t.SetOwnerID(ownerID); err != nil {
		return
	}
	t.AssigneeID = UserIDZero
	t.Status = StatusTodo
	return t, nil
}

// SetID ...
func (t *Task) SetID(id ID) error {
	if id.Equals(IDZero) {
		return common.InvalidArgumentError("id", "zero")
	}
	if !t.ID.Equals(IDZero) {
		return common.IllegalOperationError("user.id is already set")
	}
	t.ID = id
	return nil
}

// SetName ...
func (t *Task) SetName(name string) error {
	if name == "" {
		return common.InvalidArgumentError("name", "empty")
	}
	t.Name = name
	return nil
}

// SetDescription ...
func (t *Task) SetDescription(description string) error {
	if description == "" {
		return common.InvalidArgumentError("description", "empty")
	}
	t.Description = description
	return nil
}

// SetOwnerID ...
func (t *Task) SetOwnerID(ownerID UserID) error {
	if ownerID.Equals(UserIDZero) {
		return common.InvalidArgumentError("owner_id", "zero")
	}
	if !t.OwnerID.Equals(UserIDZero) {
		return common.IllegalOperationError("owner_id is already set")
	}
	t.OwnerID = ownerID
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
