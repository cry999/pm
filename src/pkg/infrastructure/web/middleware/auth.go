package middleware

import (
	"fmt"

	"github.com/cry999/pm-projects/pkg/domain/errors/common"
	"github.com/cry999/pm-projects/pkg/infrastructure/auth"
	"github.com/cry999/pm-projects/pkg/interfaces/web"
)

const (
	authHeaderKey       = "Authorization"
	okType              = "Bearer"
	authorizedUserIDKey = "authorized.user.id"
)

// NewAuthRequiredMiddleware ...
func NewAuthRequiredMiddleware(tokenizer *auth.Tokenizer) web.Middleware {
	return func(next web.HandlerFunc) web.HandlerFunc {
		return func(rc *web.RequestContext) error {
			header := rc.RequestHeader()
			typeAndToken := header.Get(authHeaderKey)
			var (
				authType  string
				authToken string
			)
			if _, err := fmt.Sscanf(typeAndToken, "%s %s", &authType, &authToken); err != nil {
				rc.JSONErrorResponse(common.NotAuthorizedError())
				rc.Logger().Error("failed to parse token: %s: %v", typeAndToken, err)
				return err
			}
			if authType != okType {
				rc.JSONErrorResponse(common.NotAuthorizedError())
				rc.Logger().Error("invalid token type: %s", authType)
				return fmt.Errorf("invalid token type: %s", authType)
			}

			userID, err := tokenizer.Restore(authToken)
			if err != nil {
				rc.JSONErrorResponse(common.NotAuthorizedError())
				rc.Logger().Error("failed to restore token: %v", err)
				return err
			}

			rc.SetParam("authorized.user.id", userID)

			return next(rc)
		}
	}
}
