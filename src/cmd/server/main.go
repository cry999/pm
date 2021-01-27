package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/cry999/pm-projects/pkg/application/iam"
	"github.com/cry999/pm-projects/pkg/application/project"
	"github.com/cry999/pm-projects/pkg/application/task"
	"github.com/cry999/pm-projects/pkg/domain/event"
	taskEvent "github.com/cry999/pm-projects/pkg/domain/event/task"
	"github.com/cry999/pm-projects/pkg/infrastructure/auth"
	"github.com/cry999/pm-projects/pkg/infrastructure/persistence"
	"github.com/cry999/pm-projects/pkg/infrastructure/pubsub"
	pubsubMiddleware "github.com/cry999/pm-projects/pkg/infrastructure/pubsub/middleware"
	"github.com/cry999/pm-projects/pkg/infrastructure/web/middleware"
	"github.com/cry999/pm-projects/pkg/interfaces/pubsub/subscribers"
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

	taskRepository := persistence.NewMySQLTaskRepository()
	taskService := task.NewApplicationService(taskRepository)
	taskHandler := handlers.NewTaskHandler(taskService)

	projectRepository := persistence.NewMySQLProjectRepository()
	projectService := project.NewApplicationService(projectRepository, taskRepository)
	projectHandler := handlers.NewProjectHandler(projectService)
	projectSubscriber := subscribers.NewProjectHandler(projectService)

	var wg sync.WaitGroup

	// * ------------ *
	// * Event PubSub *
	// * ------------ *
	eventBus, err := pubsub.NewRedisEventBus()
	if err != nil {
		exit("failed to generate event bus: %v", err)
	}
	eventBus.GlobalUse(pubsubMiddleware.Transaction)
	event.Register(eventBus)

	eventBus.Subscribe(taskEvent.Planned{}, projectSubscriber.PlannedTask)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer eventBus.Close()
		if err := eventBus.Start(context.Background()); err != nil {
			fmt.Fprintf(os.Stderr, "failed to run event bus: %v\n", err)
		}
	}()

	// * ------- *
	// * Web API *
	// * ------- *
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

	server.Route("POST", "/api/v1/projects", projectHandler.Create, authRequired)
	server.Route("GET", "/api/v1/projects", projectHandler.ListRelatedWithUser, authRequired)
	server.Route("GET", "/api/v1/projects/{project_id}", projectHandler.Retrieve, authRequired)
	server.Route("POST", "/api/v1/projects/{project_id}/tasks", projectHandler.PlanTask, authRequired)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer server.Close()
		if err := server.Run(":8080"); err != nil {
			fmt.Fprintf(os.Stderr, "server.Run: error: %v\n", err)
		}
	}()

	wg.Wait()
}
