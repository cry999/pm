// This file is auto generated

package task

import (
	errors "github.com/cry999/pm-projects/pkg/domain/errors/common"
)

var (
	// ProjectIDZero is zero-value
	ProjectIDZero = ProjectID{id: ""}
)

// ProjectID is the identity of the <no value>
type ProjectID struct {
	id string
}

// NewProjectID creates a new ProjectID instance
func NewProjectID(id string) (_ ProjectID, err error) {
	if id == "" {
		err = errors.InvalidArgumentError("id", "empty")
		return
	}
	return ProjectID{id: id}, nil
}

// MustNewProjectID creates a new ProjectID instance
func MustNewProjectID(rawid string) ProjectID {
	id, err := NewProjectID(rawid)
	if err != nil {
		panic(err)
	}
	return id
}

// Equals to other ?
func (id ProjectID) Equals(other ProjectID) bool {
	return id.id == other.id
}

// String representation
func (id ProjectID) String() string {
	return id.id
}
