package project

import "context"

// Repository ...
type Repository interface {
	NextIdentity(context.Context) (ID, error)
	Save(context.Context, *Project) error
	FindByID(context.Context, ID) (*Project, error)
	FindAllRelatedWithUser(context.Context, UserID) ([]*Project, error)
}
