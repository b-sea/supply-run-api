// Package domain defines all structs shared between domains.
package domain

import (
	"strings"
)

// ValidationError is raised when validation fails.
type ValidationError struct {
	Issues []string
}

func (e *ValidationError) Error() string {
	return "validation errors: " + strings.Join(e.Issues, ", ")
}
