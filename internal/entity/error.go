// Package entity defines all structs shared between domains.
package entity

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// ValidationError is raised when validation fails.
type ValidationError struct {
	Issues []string
}

func (e *ValidationError) Error() string {
	return "validation errors: " + strings.Join(e.Issues, ", ")
}

// NotFoundError is raised when an item cannot be found.
type NotFoundError struct {
	ID uuid.UUID
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("item not found: %s", e.ID)
}
