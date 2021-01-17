package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cry999/pm-projects/pkg/application/iam"
	"github.com/cry999/pm-projects/pkg/infrastructure/auth"
	"github.com/cry999/pm-projects/pkg/infrastructure/persistence"
	"github.com/cry999/pm-projects/pkg/infrastructure/web/middleware"
	"github.com/cry999/pm-projects/pkg/interfaces/web"
	"github.com/cry999/pm-projects/pkg/interfaces/web/api/handlers"
)

func exit(f string, a ...interface{}) {
	msg := fmt.Sprintf(f, a...)
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func main() {
	authenticator := auth.NewAuthenticator()
	tokenizer, err := auth.NewTokenizer("keys/jwt-rs256.key", "keys/jwt-rs256.pub.key")
	if err != nil {
		exit("failed to generate tokenizer: %v", err)
	}

	userRepository := persistence.NewMySQLUserRepository()
	userService := iam.NewApplicationService(userRepository, authenticator)
	userHandler := handlers.NewUserHandler(userService, tokenizer)

	server := web.NewServer()
	defer server.Close()

	authRequired := middleware.NewAuthRequiredMiddleware(tokenizer)

	server.GlobalUse(middleware.Transaction)
	server.Route("POST", "/api/v1/oauth/signin", userHandler.Signin)
	server.Route("POST", "/api/v1/oauth/signup", userHandler.Signup)

	server.Route("GET", "/api/v1/private", func(rc *web.RequestContext) error {
		userID, ok := rc.GetParam("authorized.user.id")
		if !ok {
			rc.JSONErrorResponse(fmt.Errorf("failed to get userID"))
			return nil
		}
		rc.JSONResponse(http.StatusOK, map[string]string{
			"message": fmt.Sprintf("Hello, %s", userID),
		})
		return nil
	}, authRequired)

	if err := server.Run(":8080"); err != nil {
		exit("server.Run: error: %v", err)
	}
}
