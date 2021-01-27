// This file is auto generated

package project

import (
	errors "github.com/cry999/pm-projects/pkg/domain/errors/common"
)

var (
	// IDZero is zero-value
	IDZero = ID{id: ""}
)

// ID is the identity of the <no value>
type ID struct {
	id string
}

// NewID creates a new ID instance
func NewID(id string) (_ ID, err error) {
	if id == "" {
		err = errors.InvalidArgumentError("id", "empty")
		return
	}
	return ID{id: id}, nil
}

// MustNewID creates a new ID instance
func MustNewID(rawid string) ID {
	id, err := NewID(rawid)
	if err != nil {
		panic(err)
	}
	return id
}

// Equals to other ?
func (id ID) Equals(other ID) bool {
	return id.id == other.id
}

// String representation
func (id ID) String() string {
	return id.id
}
