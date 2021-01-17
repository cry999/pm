package commands

import "github.com/cry999/pm-projects/pkg/domain/model/iam"

// UserDescriptor ...
type UserDescriptor struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// NewUserDescriptor creates a new UserDescriptor instance from domain.User
func NewUserDescriptor(u *iam.User) UserDescriptor {
	return UserDescriptor{
		ID:    u.ID.String(),
		Email: u.Email,
	}
}
