package main

import (
	"fmt"
	"os"

	"github.com/cry999/pm-projects/pkg/application/iam"
	"github.com/cry999/pm-projects/pkg/application/task"
	"github.com/cry999/pm-projects/pkg/infrastructure/auth"
	"github.com/cry999/pm-projects/pkg/infrastructure/persistence"
	"github.com/cry999/pm-projects/pkg/infrastructure/web/middleware"
	"github.com/cry999/pm-projects/pkg/interfaces/web"
	"github.com/cry999/pm-projects/pkg/interfaces/web/api/handlers"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func exit(f string, a ...interface{}) {
	msg := fmt.Sprintf(f, a...)
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func main() {
	boil.DebugMode = true

	authenticator := auth.NewAuthenticator()
	tokenizer, err := auth.NewTokenizer("keys/jwt-rs256.key", "keys/jwt-rs256.pub.key")
	if err != nil {
		exit("failed to generate tokenizer: %v", err)
	}

	userRepository := persistence.NewMySQLUserRepository()
	userService := iam.NewApplicationService(userRepository, authenticator)
	userHandler := handlers.NewUserHandler(userService, tokenizer)

	taskRepository := persistence.NewMySQLTaskRepository()
	taskService := task.NewApplicationService(taskRepository)
	taskHandler := handlers.NewTaskHandler(taskService)

	server := web.NewServer()
	defer server.Close()

	authRequired := middleware.NewAuthRequiredMiddleware(tokenizer)

	server.GlobalUse(middleware.Transaction)
	server.Route("POST", "/api/v1/oauth/signin", userHandler.Signin)
	server.Route("POST", "/api/v1/oauth/signup", userHandler.Signup)

	server.Route("GET", "/api/v1/tasks", taskHandler.ListAssociatedWithUserTasks, authRequired)
	server.Route("POST", "/api/v1/tasks", taskHandler.CreateTask, authRequired)
	server.Route("GET", "/api/v1/tasks/{task_id}", taskHandler.RetrieveTask, authRequired)
	server.Route("PUT", "/api/v1/tasks/{task_id}/{action:progress|done|redo|cancel|hold}", taskHandler.UpdateStatus, authRequired)
	server.Route("PUT", "/api/v1/tasks/{task_id}/assign", taskHandler.AssignSignedInUserToTask, authRequired)

	if err := server.Run(":8080"); err != nil {
		exit("server.Run: error: %v", err)
	}
}
