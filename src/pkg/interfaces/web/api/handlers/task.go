package handlers

import (
	"net/http"

	"github.com/cry999/pm-projects/pkg/application/task"
	"github.com/cry999/pm-projects/pkg/application/task/commands"
	"github.com/cry999/pm-projects/pkg/interfaces/web"
	"github.com/cry999/pm-projects/pkg/interfaces/web/api/errors"
)

// TaskHandler ...
type TaskHandler struct {
	service task.ApplicationService
}

// NewTaskHandler creates a new TaskHandler instance
func NewTaskHandler(service task.ApplicationService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

// CreateTask ...
func (h *TaskHandler) CreateTask(rc *web.RequestContext) (err error) {
	userID, err := rc.GetParamString("authorized.user.id")
	if err != nil {
		err = errors.HTTPErrorf(http.StatusUnauthorized, "not authorized")
		rc.JSONErrorResponse(err)
		rc.Logger().Error("cannot get authorized user id: %v", err)
		return
	}

	var input commands.CreateTaskInput
	if err = rc.JSONRequest(&input); err != nil {
		err = errors.HTTPErrorf(http.StatusBadRequest, "invalid json: %v", err)
		rc.JSONErrorResponse(err)
		return
	}
	input.OwnerID = userID

	output, err := h.service.CreateTask(rc.Context(), input)
	if err != nil {
		rc.JSONErrorResponse(err)
		return
	}
	rc.JSONResponse(http.StatusCreated, output)
	return
}

// RetrieveTask ...
func (h *TaskHandler) RetrieveTask(rc *web.RequestContext) (err error) {
	userID, err := rc.GetParamString("authorized.user.id")
	if err != nil {
		err = errors.HTTPErrorf(http.StatusUnauthorized, "not authorized")
		rc.JSONErrorResponse(err)
		rc.Logger().Error("cannot get authorized user id: %v", err)
		return
	}
	taskID, err := rc.GetParamString("path.task_id")
	if err != nil {
		err = errors.HTTPErrorf(http.StatusNotFound, "404 page not found")
		rc.JSONErrorResponse(err)
		rc.Logger().Error("cannot get task_id: %v", err)
		return
	}
	input := commands.RetrieveTaskInput{
		UserID: userID,
		TaskID: taskID,
	}
	output, err := h.service.RetrieveTask(rc.Context(), input)
	if err != nil {
		rc.JSONErrorResponse(err)
		return
	}
	rc.JSONResponse(http.StatusOK, output)
	return
}

// UpdateStatus ...
func (h *TaskHandler) UpdateStatus(rc *web.RequestContext) (err error) {
	actorID, err := rc.GetParamString("authorized.user.id")
	if err != nil {
		rc.Logger().Error("cannot get authorized user id: %v", err)
		err = errors.HTTPErrorf(http.StatusUnauthorized, "not authorized")
		rc.JSONErrorResponse(err)
		return
	}
	taskID, err := rc.GetParamString("path.task_id")
	if err != nil {
		rc.Logger().Error("cannot get task_id: %v", err)
		err = errors.HTTPErrorf(http.StatusNotFound, "404 page not found")
		rc.JSONErrorResponse(err)
		return
	}
	action, err := rc.GetParamString("path.action")
	if err != nil {
		rc.Logger().Error("cannot get action: %v", err)
		err = errors.HTTPErrorf(http.StatusNotFound, "404 page not found")
		rc.JSONErrorResponse(err)
		return
	}
	available := false
	for _, defined := range commands.Actions {
		if action == defined {
			available = true
			break
		}
	}
	if !available {
		err = errors.HTTPErrorf(http.StatusNotFound, "404 page not found: %s", action)
		rc.JSONErrorResponse(err)
		return
	}

	input := commands.UpdateStatusInput{
		ActorID: actorID,
		TaskID:  taskID,
		Action:  action,
	}
	output, err := h.service.UpdateStatus(rc.Context(), input)
	if err != nil {
		rc.JSONErrorResponse(err)
		return
	}
	rc.JSONResponse(http.StatusOK, output)
	return
}

// ListAssociatedWithUserTasks ...
func (h *TaskHandler) ListAssociatedWithUserTasks(rc *web.RequestContext) (err error) {
	userID, err := rc.GetParamString("authorized.user.id")
	if err != nil {
		rc.JSONErrorResponse(err)
		err = errors.HTTPErrorf(http.StatusUnauthorized, "not authorized")
		rc.Logger().Error("cannot get authorized user id: %v", err)
		return
	}

	input := commands.GetAssociatedWithUserTasksInput{
		UserID: userID,
	}
	output, err := h.service.GetAssociatedWithUserTasks(rc.Context(), input)
	if err != nil {
		rc.JSONErrorResponse(err)
		return
	}
	rc.JSONResponse(http.StatusOK, output)
	return
}

// AssignSignedInUserToTask ...
func (h *TaskHandler) AssignSignedInUserToTask(rc *web.RequestContext) (err error) {
	userID, err := rc.GetParamString("authorized.user.id")
	if err != nil {
		rc.JSONErrorResponse(err)
		err = errors.HTTPErrorf(http.StatusUnauthorized, "not authorized")
		rc.Logger().Error("cannot get authorized user id: %v", err)
		return
	}
	taskID, err := rc.GetParamString("path.task_id")
	if err != nil {
		rc.JSONErrorResponse(err)
		err = errors.HTTPErrorf(http.StatusUnauthorized, "not authorized")
		rc.Logger().Error("cannot get authorized user id: %v", err)
		return
	}
	input := commands.AssignUserToTaskInput{
		AssigneeID: userID,
		TaskID:     taskID,
	}
	output, err := h.service.AssignUserToTask(rc.Context(), input)
	if err != nil {
		rc.JSONErrorResponse(err)
		return
	}
	rc.JSONResponse(http.StatusOK, output)
	return
}
