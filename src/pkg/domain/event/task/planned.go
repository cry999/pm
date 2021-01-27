package task

import "time"

// Planned ...
type Planned struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	OwnerID     string     `json:"owner_id"`
	ProjectID   string     `json:"project_id"`
	AssigneeID  string     `json:"assignee_id"`
	Status      string     `json:"status"`
	Deadline    *time.Time `json:"deadline"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// Type of the Event
func (Planned) Type() string { return "pm.projects.task.Planned" }
