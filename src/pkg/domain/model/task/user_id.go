// This file is auto generated

package task

import (
	errors "github.com/cry999/pm-projects/pkg/domain/errors/common"
)

var (
	// UserIDZero is zero-value
	UserIDZero = UserID{id: ""}
)

// UserID is the identity of the <no value>
type UserID struct {
	id string
}

// NewUserID creates a new UserID instance
func NewUserID(id string) (_ UserID, err error) {
	if id == "" {
		err = errors.InvalidArgumentError("id", "empty")
		return
	}
	return UserID{id: id}, nil
}

// MustNewUserID creates a new UserID instance
func MustNewUserID(rawid string) UserID {
	id, err := NewUserID(rawid)
	if err != nil {
		panic(err)
	}
	return id
}

// Equals to other ?
func (id UserID) Equals(other UserID) bool {
	return id.id == other.id
}

// String representation
func (id UserID) String() string {
	return id.id
}
