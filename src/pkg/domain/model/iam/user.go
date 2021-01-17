package iam

import (
	"regexp"

	"github.com/cry999/pm-projects/pkg/domain/errors/common"
)

var (
	emailRegexp = regexp.MustCompile(`\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
)

// User ...
type User struct {
	ID             UserID
	Email          string
	HashedPassword string
}

// NewUser creates a new User instance
func NewUser(id UserID, email, hashedPassword string) (_ *User, err error) {
	u := new(User)
	if err = u.SetID(id); err != nil {
		return
	}
	if err = u.SetEmail(email); err != nil {
		return
	}
	if err = u.SetHashedPassword(hashedPassword); err != nil {
		return
	}
	return u, nil
}

// SetID ...
func (u *User) SetID(id UserID) error {
	if id.Equals(UserIDZero) {
		return common.InvalidArgumentError("id", "empty")
	}
	if !u.ID.Equals(UserIDZero) {
		return common.IllegalOperationError("user id is already set")
	}
	u.ID = id
	return nil
}

// SetEmail ...
func (u *User) SetEmail(email string) error {
	if !emailRegexp.MatchString(email) {
		return common.InvalidArgumentError("email", "invalid format")
	}
	u.Email = email
	return nil
}

// SetHashedPassword ...
func (u *User) SetHashedPassword(hashedPassword string) error {
	if hashedPassword == "" {
		return common.InvalidArgumentError("hashedPassword", "empty")
	}
	u.HashedPassword = hashedPassword
	return nil
}
