package entity

import (
	"errors"
	"strings"
)

// ErrNotFound is raised when an entity cannot be found.
var ErrNotFound = errors.New("not found")

// ValidationError is holds validation errors.
type ValidationError struct {
	InnerErrors []error
}

func (e *ValidationError) Error() string {
	messages := make([]string, len(e.InnerErrors))

	for i := range e.InnerErrors {
		messages[i] = e.InnerErrors[i].Error()
	}

	return "validation errors: " + strings.Join(messages, ", ")
}

// IsEmpty returns if the error has any inner errors.
func (e *ValidationError) IsEmpty() bool {
	return len(e.InnerErrors) == 0
}
