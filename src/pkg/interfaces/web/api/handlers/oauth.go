package handlers

import (
	"net/http"

	"github.com/cry999/pm-projects/pkg/application/iam"
	"github.com/cry999/pm-projects/pkg/application/iam/commands"
	"github.com/cry999/pm-projects/pkg/interfaces/web"
	"github.com/cry999/pm-projects/pkg/interfaces/web/api/errors"
)

// Tokenizer ...
type Tokenizer interface {
	Tokenize(userID string) (token string, err error)
}

// UserHandler ...
type UserHandler struct {
	service   iam.ApplicationService
	tokenizer Tokenizer
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(
	service iam.ApplicationService,
	tokenizer Tokenizer,
) *UserHandler {
	return &UserHandler{
		service:   service,
		tokenizer: tokenizer,
	}
}

// Signin ...
func (h *UserHandler) Signin(ctx *web.RequestContext) error {
	var input commands.SigninInput
	if err := ctx.JSONRequest(&input); err != nil {
		ctx.JSONErrorResponse(
			errors.HTTPErrorf(http.StatusBadRequest, "invalid json: %v", err),
		)
		return err
	}

	output, err := h.service.Signin(ctx.Context(), input)
	if err != nil {
		ctx.JSONErrorResponse(err)
		return err
	}

	token, err := h.tokenizer.Tokenize(output.ID)
	if err != nil {
		ctx.JSONErrorResponse(err)
		return err
	}

	ctx.JSONResponse(http.StatusOK, map[string]string{
		"access_token": token,
	})
	return nil
}

// Signup ...
func (h *UserHandler) Signup(ctx *web.RequestContext) (err error) {
	var input commands.SignupInput
	if err = ctx.JSONRequest(&input); err != nil {
		ctx.JSONErrorResponse(
			errors.HTTPErrorf(http.StatusBadRequest, "invalid json: %v", err),
		)
		return
	}

	output, err := h.service.Signup(ctx.Context(), input)
	if err != nil {
		ctx.JSONErrorResponse(err)
		return
	}

	token, err := h.tokenizer.Tokenize(output.ID)
	if err != nil {
		ctx.JSONErrorResponse(err)
		return err
	}

	ctx.JSONResponse(http.StatusOK, map[string]string{
		"access_token": token,
	})
	return
}
