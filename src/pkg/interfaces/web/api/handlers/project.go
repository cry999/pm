package handlers

import (
	"net/http"

	"github.com/cry999/pm-projects/pkg/application/project"
	"github.com/cry999/pm-projects/pkg/application/project/commands"
	"github.com/cry999/pm-projects/pkg/interfaces/web"
	"github.com/cry999/pm-projects/pkg/interfaces/web/api/errors"
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

// Create ...
func (h *ProjectHandler) Create(rc *web.RequestContext) (err error) {
	var input commands.CreateProjectInput
	if err := rc.JSONRequest(&input); err != nil {
		rc.JSONErrorResponse(errors.HTTPErrorf(http.StatusBadRequest, "invalid json"))
		return err
	}
	userID, err := rc.GetParamString("authorized.user.id")
	if err != nil {
		rc.JSONErrorResponse(errors.HTTPErrorf(http.StatusUnauthorized, "not authorized"))
		return err
	}
	input.OwnerID = userID

	output, err := h.service.CreateProject(rc.Context(), input)
	if err != nil {
		rc.JSONErrorResponse(err)
		return err
	}

	rc.JSONResponse(http.StatusCreated, output)
	return
}

// Retrieve ...
func (h *ProjectHandler) Retrieve(rc *web.RequestContext) (err error) {
	userID, err := rc.GetParamString("authorized.user.id")
	if err != nil {
		rc.JSONErrorResponse(errors.HTTPErrorf(http.StatusUnauthorized, "not authorized"))
		return
	}
	projectID, err := rc.GetParamString("path.project_id")
	if err != nil {
		rc.JSONErrorResponse(errors.HTTPErrorf(http.StatusNotFound, "404 page not found"))
		return
	}
	input := commands.RetrieveProjectInput{
		OwnerID:   userID,
		ProjectID: projectID,
	}
	output, err := h.service.RetrieveProject(rc.Context(), input)
	if err != nil {
		rc.JSONErrorResponse(err)
		return
	}
	rc.JSONResponse(http.StatusOK, output)
	return
}

// ListRelatedWithUser ...
func (h *ProjectHandler) ListRelatedWithUser(rc *web.RequestContext) (err error) {
	userID, err := rc.GetParamString("authorized.user.id")
	if err != nil {
		rc.JSONErrorResponse(errors.HTTPErrorf(http.StatusUnauthorized, "not authorized"))
		return
	}
	input := commands.ListProjectsInput{
		UserID: userID,
	}
	output, err := h.service.ListProjects(rc.Context(), input)
	if err != nil {
		rc.JSONErrorResponse(err)
		return
	}
	rc.JSONResponse(http.StatusOK, output)
	return
}

// PlanTask ...
func (h *ProjectHandler) PlanTask(rc *web.RequestContext) (err error) {
	userID, err := rc.GetParamString("authorized.user.id")
	if err != nil {
		rc.JSONErrorResponse(errors.HTTPErrorf(http.StatusUnauthorized, "not authorized"))
		return
	}
	projectID, err := rc.GetParamString("path.project_id")
	var input commands.PlanTaskInput
	if err = rc.JSONRequest(&input); err != nil {
		rc.JSONErrorResponse(errors.HTTPErrorf(http.StatusBadRequest, "invalid json"))
		rc.Logger().Error("failed to decode json body: %v", err)
		return
	}
	input.OwnerID = userID
	input.ProjectID = projectID

	output, err := h.service.PlanTask(rc.Context(), input)
	if err != nil {
		rc.JSONErrorResponse(err)
		return
	}
	rc.JSONResponse(http.StatusCreated, output)
	return
}
