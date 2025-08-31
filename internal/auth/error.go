// Package auth defines authentication workflows and entities.
package auth

import (
	"errors"
)

// ErrUnauthorized, et al. are custom auth-based errors.
var (
	ErrUnauthorized     = errors.New("unauthorized")
	ErrNotFound         = errors.New("entity not found")
	ErrDuplicateAccount = errors.New("duplicate user account")
)
