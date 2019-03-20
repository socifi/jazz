package config

import (
	"fmt"
)

type TraverseError struct {
	Cause string
	Key   string
}

func (e *TraverseError) Error() string {
	return fmt.Sprintf("%v: %v", e.Cause, e.Key)
}

func NewTraverseError(cause, key string) *TraverseError {
	return &TraverseError{cause, key}
}
