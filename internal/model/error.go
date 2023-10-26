// Package model defines all data entities shared between front end, service, and repository layers.
package model

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Issues []string `json:"issues"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation errors: %s", strings.Join(e.Issues, ", "))
}

type NotFoundError struct {
	ID ID `json:"id"`
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("not found: %s", e.ID)
}
