package project

import (
	"context"

	"github.com/cry999/pm-projects/pkg/application/project/commands"
	taskCommands "github.com/cry999/pm-projects/pkg/application/task/commands"
	"github.com/cry999/pm-projects/pkg/domain/errors/common"
	"github.com/cry999/pm-projects/pkg/domain/model/project"
	"github.com/cry999/pm-projects/pkg/domain/model/task"
)

// ApplicationService ...
type ApplicationService interface {
	CreateProject(context.Context, commands.CreateProjectInput) (*commands.CreateProjectOutput, error)
	RetrieveProject(context.Context, commands.RetrieveProjectInput) (*commands.RetrieveProjectOutput, error)
	ListProjects(context.Context, commands.ListProjectsInput) (*commands.ListProjectsOutput, error)
	PlanTask(context.Context, commands.PlanTaskInput) (*commands.PlanTaskOutput, error)
	PlannedTask(context.Context, commands.PlannedTaskInput) (*commands.PlannedTaskOutput, error)
}

// service ...
type service struct {
	projectRepository project.Repository
	taskRepository    task.Repository
}

// NewApplicationService ...
func NewApplicationService(
	projectRepository project.Repository,
	taskRepository task.Repository,
) ApplicationService {
	return &service{
		projectRepository: projectRepository,
		taskRepository:    taskRepository,
	}
}

// CreateProject ...
func (s *service) CreateProject(ctx context.Context, input commands.CreateProjectInput) (_ *commands.CreateProjectOutput, err error) {
	id, err := s.projectRepository.NextIdentity(ctx)
	if err != nil {
		return
	}
	ownerID, err := project.NewUserID(input.OwnerID)
	if err != nil {
		return
	}
	prj, err := project.New(id, ownerID, input.Name, input.ElevatorPitch)
	if err != nil {
		return
	}
	if err = s.projectRepository.Save(ctx, prj); err != nil {
		return
	}
	return &commands.CreateProjectOutput{
		ProjectDescriptor: commands.NewProjectDescriptor(prj, nil),
	}, nil
}

func (s *service) RetrieveProject(ctx context.Context, input commands.RetrieveProjectInput) (_ *commands.RetrieveProjectOutput, err error) {
	ownerID, err := project.NewUserID(input.OwnerID)
	if err != nil {
		return
	}
	projectID, err := project.NewID(input.ProjectID)
	if err != nil {
		return
	}
	prj, err := s.projectRepository.FindByID(ctx, projectID)
	if err != nil {
		return
	}
	// TODO: アクセス可能か判定するメソッドを Project に持たせる
	if !prj.OwnerID.Equals(ownerID) {
		return nil, common.ForbiddenError(ownerID.String(), "read.proejct")
	}
	tasks, err := s.getTasks(ctx, prj)
	if err != nil {
		return
	}
	return &commands.RetrieveProjectOutput{
		ProjectDescriptor: commands.NewProjectDescriptor(prj, tasks),
	}, nil
}

func (s *service) ListProjects(ctx context.Context, input commands.ListProjectsInput) (_ *commands.ListProjectsOutput, err error) {
	userID, err := project.NewUserID(input.UserID)
	if err != nil {
		return
	}
	projects, err := s.projectRepository.FindAllRelatedWithUser(ctx, userID)
	if err != nil {
		return
	}
	out := &commands.ListProjectsOutput{
		Results: []commands.ProjectDescriptor{},
	}
	for _, prj := range projects {
		tasks, err := s.getTasks(ctx, prj)
		if err != nil {
			return nil, err
		}
		out.Results = append(out.Results, commands.NewProjectDescriptor(prj, tasks))
	}
	return out, nil
}

func (s *service) PlanTask(ctx context.Context, input commands.PlanTaskInput) (_ *commands.PlanTaskOutput, err error) {
	ownerID, err := task.NewUserID(input.OwnerID)
	if err != nil {
		return
	}
	projectID, err := project.NewID(input.ProjectID)
	if err != nil {
		return
	}
	taskID, err := s.taskRepository.NextIdentity(ctx)
	if err != nil {
		return
	}
	project, err := s.projectRepository.FindByID(ctx, projectID)
	if err != nil {
		return
	}
	plannedTask, err := project.PlanTask(taskID, input.Name, input.Description, ownerID)
	if err != nil {
		return
	}
	if err = s.taskRepository.Save(ctx, plannedTask); err != nil {
		return
	}
	return &commands.PlanTaskOutput{
		TaskDescriptor: taskCommands.NewTaskDescriptor(plannedTask),
	}, nil
}

func (s *service) PlannedTask(ctx context.Context, input commands.PlannedTaskInput) (_ *commands.PlannedTaskOutput, err error) {
	projectID, err := project.NewID(input.ProjectID)
	if err != nil {
		return
	}
	plannedTaskID, err := project.NewPlannedTaskID(input.PlannedTaskID)
	if err != nil {
		return
	}
	prj, err := s.projectRepository.FindByID(ctx, projectID)
	if err != nil {
		return
	}
	if err = prj.PlannedTask(plannedTaskID); err != nil {
		return
	}
	if err = s.projectRepository.Save(ctx, prj); err != nil {
		return
	}
	tasks, err := s.getTasks(ctx, prj)
	if err != nil {
		return
	}
	return &commands.PlannedTaskOutput{
		ProjectDescriptor: commands.NewProjectDescriptor(prj, tasks),
	}, nil
}

// TODO: taskRepository にアクセスしているのをどう分離していくかを考える必要あり
func (s *service) getTasks(ctx context.Context, prj *project.Project) ([]*task.Task, error) {
	tasks := []*task.Task{}
	for _, plannedTaskID := range prj.PlannedTasks {
		taskID, err := task.NewID(plannedTaskID.String())
		if err != nil {
			return nil, err
		}
		task, err := s.taskRepository.Find(ctx, taskID)
		if err != nil {
			// TODO: getTasks may fail because `Planned` event may be thrown before completion
			// TODO: of creating task. So, currently error is ignored
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
