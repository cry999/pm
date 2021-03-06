// This file is auto generated by go-ddd

package iam

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

// Equals to other ?
func (id UserID) Equals(other UserID) bool {
	return id.id == other.id
}

// String representation
func (id UserID) String() string {
	return id.id
}
