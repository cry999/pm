package subscribers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/cry999/pm-projects/pkg/application/project"
	"github.com/cry999/pm-projects/pkg/application/project/commands"
	"github.com/cry999/pm-projects/pkg/domain/event/task"
)

// ProjectHandler ...
type ProjectHandler struct {
	service project.ApplicationService
}

// NewProjectHandler creates a new ProjectHandler instance
func NewProjectHandler(service project.ApplicationService) *ProjectHandler {
	return &ProjectHandler{
		service: service,
	}
}

// PlannedTask ...
func (h *ProjectHandler) PlannedTask(ctx context.Context, payload io.Reader) error {
	var event task.Planned
	if err := json.NewDecoder(payload).Decode(&event); err != nil {
		return fmt.Errorf("failed to decode payload: %v", err)
	}
	input := commands.PlannedTaskInput{
		ProjectID:     event.ProjectID,
		PlannedTaskID: event.ID,
	}
	if _, err := h.service.PlannedTask(ctx, input); err != nil {
		return fmt.Errorf("project.ApplicationService.PlannedTask: error: %v", err)
	}
	return nil
}
