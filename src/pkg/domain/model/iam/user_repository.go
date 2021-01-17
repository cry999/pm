package iam

import "context"

// UserRepository ...
type UserRepository interface {
	NextIdentity(context.Context) (UserID, error)
	Save(context.Context, *User) error
	Find(context.Context, UserID) (*User, error)
	FindByEmail(context.Context, string) (*User, error)
}
