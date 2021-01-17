package iam

import (
	"context"

	"github.com/cry999/pm-projects/pkg/application/iam/commands"
	"github.com/cry999/pm-projects/pkg/domain/model/iam"
)

// ApplicationService ...
type ApplicationService interface {
	Signin(context.Context, commands.SigninInput) (*commands.SigninOutput, error)
	Signup(context.Context, commands.SignupInput) (*commands.SignupOutput, error)
}

// service ...
type service struct {
	repository    iam.UserRepository
	authenticator iam.Authenticator
}

// NewApplicationService creates a new ApplicationService instance
func NewApplicationService(
	repository iam.UserRepository,
	authenticator iam.Authenticator,
) ApplicationService {
	return &service{
		repository:    repository,
		authenticator: authenticator,
	}
}

// Signin ...
func (s *service) Signin(ctx context.Context, input commands.SigninInput) (_ *commands.SigninOutput, err error) {
	user, err := s.repository.FindByEmail(ctx, input.Email)
	if err != nil {
		return
	}
	if err = s.authenticator.Auth(input.Password, user.HashedPassword); err != nil {
		return
	}
	return &commands.SigninOutput{
		UserDescriptor: commands.NewUserDescriptor(user),
	}, nil
}

// Signup ...
func (s *service) Signup(ctx context.Context, input commands.SignupInput) (_ *commands.SignupOutput, err error) {
	userID, err := s.repository.NextIdentity(ctx)
	if err != nil {
		return
	}
	hashed, err := s.authenticator.Hash(input.Password)
	if err != nil {
		return
	}
	user, err := iam.NewUser(userID, input.Email, hashed)
	if err != nil {
		return
	}
	if err = s.repository.Save(ctx, user); err != nil {
		return
	}
	return &commands.SignupOutput{
		UserDescriptor: commands.NewUserDescriptor(user),
	}, nil
}
