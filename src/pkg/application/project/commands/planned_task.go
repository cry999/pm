package commands

// TODO: task domain に入れるべきか要検討

// PlannedTaskInput ...
type PlannedTaskInput struct {
	ProjectID     string `json:"project_id"`
	PlannedTaskID string `json:"planned_task_id"`
}

// PlannedTaskOutput ...
type PlannedTaskOutput struct {
	ProjectDescriptor
}
