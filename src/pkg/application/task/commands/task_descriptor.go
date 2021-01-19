package commands

import (
	"time"

	"github.com/cry999/pm-projects/pkg/domain/model/task"
)

// TaskDescriptor ...
type TaskDescriptor struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	OwnerID     *string    `json:"owner_id"`
	AssigneeID  *string    `json:"assignee_id"`
	Deadline    *time.Time `json:"deadline"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// NewTaskDescriptor creates a new TaskDescriptor instance
func NewTaskDescriptor(task *task.Task) TaskDescriptor {
	return TaskDescriptor{
		ID:          task.ID.String(),
		Name:        task.Name,
		Description: task.Description,
		Status:      string(task.Status),
		OwnerID:     nilIfEmpty(task.OwnerID.String()),
		AssigneeID:  nilIfEmpty(task.AssigneeID.String()),
		Deadline:    task.Deadline,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func nilIfEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
