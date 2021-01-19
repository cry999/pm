package web

import (
	"net/http"

	"github.com/cry999/pm-projects/pkg/domain/errors/common"
	"github.com/cry999/pm-projects/pkg/interfaces/web/api/errors"
)

// ErrorCode ...
func ErrorCode(err error) (code int) {
	switch err := err.(type) {
	case *common.InvalidArgument:
		return http.StatusBadRequest // 400
	case *common.IllegalOperation:
		return http.StatusUnprocessableEntity // 422
	case *common.NotFound:
		return http.StatusNotFound
	case *common.NotAuthorized:
		return http.StatusUnauthorized
	case *common.Forbidden:
		return http.StatusForbidden
	case *errors.HTTPError:
		return err.Code
	}
	return http.StatusInternalServerError
}
