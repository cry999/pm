package task

import (
	"context"

	"github.com/cry999/pm-projects/pkg/application/task/commands"
	"github.com/cry999/pm-projects/pkg/domain/errors/common"
	domain "github.com/cry999/pm-projects/pkg/domain/model/task"
)

// ApplicationService ...
type ApplicationService interface {
	CreateTask(context.Context, commands.CreateTaskInput) (*commands.CreateTaskOutput, error)
	RetrieveTask(context.Context, commands.RetrieveTaskInput) (*commands.RetrieveTaskOutput, error)
	UpdateStatus(context.Context, commands.UpdateStatusInput) (*commands.UpdateStatusOutput, error)
	GetAssociatedWithUserTasks(context.Context, commands.GetAssociatedWithUserTasksInput) (*commands.GetAssociatedWithUserTasksOutput, error)
	AssignUserToTask(context.Context, commands.AssignUserToTaskInput) (*commands.AssignUserToTaskOutput, error)
}

type service struct {
	repository domain.Repository
}

// NewApplicationService creates a new ApplicationService implements
func NewApplicationService(repository domain.Repository) ApplicationService {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateTask(ctx context.Context, input commands.CreateTaskInput) (output *commands.CreateTaskOutput, err error) {
	ownerID, err := domain.NewUserID(input.OwnerID)
	if err != nil {
		return
	}
	taskID, err := s.repository.NextIdentity(ctx)
	if err != nil {
		return
	}
	task, err := domain.NewTask(taskID, input.Name, input.Description, ownerID)
	if err != nil {
		return
	}
	if err = s.repository.Save(ctx, task); err != nil {
		return
	}
	return &commands.CreateTaskOutput{
		TaskDescriptor: commands.NewTaskDescriptor(task),
	}, nil
}

func (s *service) RetrieveTask(ctx context.Context, input commands.RetrieveTaskInput) (_ *commands.RetrieveTaskOutput, err error) {
	userID, err := domain.NewUserID(input.UserID)
	if err != nil {
		return
	}
	taskID, err := domain.NewID(input.TaskID)
	if err != nil {
		return
	}
	task, err := s.repository.Find(ctx, taskID)
	if err != nil {
		return
	}
	if !task.CanAccess(userID) {
		err = common.ForbiddenError(userID.String(), "task.read")
		return
	}
	return &commands.RetrieveTaskOutput{
		TaskDescriptor: commands.NewTaskDescriptor(task),
	}, nil
}

func (s *service) UpdateStatus(ctx context.Context, input commands.UpdateStatusInput) (_ *commands.UpdateStatusOutput, err error) {
	actorID, err := domain.NewUserID(input.ActorID)
	if err != nil {
		return
	}
	taskID, err := domain.NewID(input.TaskID)
	if err != nil {
		return
	}
	task, err := s.repository.Find(ctx, taskID)
	if err != nil {
		return
	}
	// TODO: CanAccess に error を返させた方が良さそう
	// ? そもそも、各メソッドに actorID も渡して各メソッドで権限あるかを確認した方が良さそう。
	// ? AccessRole クラスを task context に作る?
	if !task.CanAccess(actorID) {
		err = common.ForbiddenError(actorID.String(), "task.update_status")
		return
	}
	switch input.Action {
	case "progress":
		err = task.Progress()
	case "done":
		err = task.Done()
	case "redo":
		err = task.Redo()
	case "cancel":
		err = task.Cancel()
	case "hold":
		err = task.PutOnHold()
	default:
		err = common.IllegalOperationError("invalid action: %s", input.Action)
	}
	if err != nil {
		return
	}
	if err = s.repository.Save(ctx, task); err != nil {
		return
	}
	return &commands.UpdateStatusOutput{
		TaskDescriptor: commands.NewTaskDescriptor(task),
	}, nil
}

func (s *service) GetAssociatedWithUserTasks(ctx context.Context, input commands.GetAssociatedWithUserTasksInput) (out *commands.GetAssociatedWithUserTasksOutput, err error) {
	userID, err := domain.NewUserID(input.UserID)
	if err != nil {
		return
	}
	tasks, err := s.repository.FindAllAssociatedWithUser(ctx, userID)
	if err != nil {
		return
	}
	out = &commands.GetAssociatedWithUserTasksOutput{
		Results: []commands.TaskDescriptor{},
	}
	for _, task := range tasks {
		out.Results = append(out.Results, commands.NewTaskDescriptor(task))
	}
	return
}

func (s *service) AssignUserToTask(ctx context.Context, input commands.AssignUserToTaskInput) (output *commands.AssignUserToTaskOutput, err error) {
	assigneeID, err := domain.NewUserID(input.AssigneeID)
	if err != nil {
		return
	}
	taskID, err := domain.NewID(input.TaskID)
	if err != nil {
		return
	}
	task, err := s.repository.Find(ctx, taskID)
	if err != nil {
		return
	}
	if err = task.Assign(assigneeID); err != nil {
		return
	}
	if err = s.repository.Save(ctx, task); err != nil {
		return
	}
	return &commands.AssignUserToTaskOutput{
		TaskDescriptor: commands.NewTaskDescriptor(task),
	}, nil
}
