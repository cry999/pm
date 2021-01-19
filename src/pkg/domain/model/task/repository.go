package task

import "context"

// Repository ...
type Repository interface {
	NextIdentity(context.Context) (ID, error)
	Save(context.Context, *Task) error
	Find(context.Context, ID) (*Task, error)
	FindAllAssociatedWithUser(context.Context, UserID) ([]*Task, error)
	// TODO: 以下のメソッドは直近でいらない
	// FindAllOwnedTask(context.Context, UserID) ([]*Task, error)
	// FindAllAssignedTask(context.Context, UserID) ([]*Task, error)
}
