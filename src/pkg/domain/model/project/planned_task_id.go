// This file is auto generated

package project

import (
	errors "github.com/cry999/pm-projects/pkg/domain/errors/common"
)

var (
	// PlannedTaskIDZero is zero-value
	PlannedTaskIDZero = PlannedTaskID{id: ""}
)

// PlannedTaskID is the identity of the <no value>
type PlannedTaskID struct {
	id string
}

// NewPlannedTaskID creates a new PlannedTaskID instance
func NewPlannedTaskID(id string) (_ PlannedTaskID, err error) {
	if id == "" {
		err = errors.InvalidArgumentError("id", "empty")
		return
	}
	return PlannedTaskID{id: id}, nil
}

// MustNewPlannedTaskID creates a new PlannedTaskID instance
func MustNewPlannedTaskID(rawid string) PlannedTaskID {
	id, err := NewPlannedTaskID(rawid)
	if err != nil {
		panic(err)
	}
	return id
}

// Equals to other ?
func (id PlannedTaskID) Equals(other PlannedTaskID) bool {
	return id.id == other.id
}

// String representation
func (id PlannedTaskID) String() string {
	return id.id
}
