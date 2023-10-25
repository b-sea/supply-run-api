// Package service implements all business logic for the API.
package service

import (
	"fmt"
	"strings"

	"github.com/b-sea/supply-run-api/internal/model"
)

var (
	ErrAuthentication = fmt.Errorf("authentication error")
	ErrAuthorization  = fmt.Errorf("authorization error")
)

type ValidationError struct {
	Issues []string `json:"issues"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation errors: %s", strings.Join(e.Issues, ", "))
}

type NotFoundError struct {
	ID model.ID `json:"id"`
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("not found: %s", e.ID)
}
