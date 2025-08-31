// Package auth implements authentication workflows and entities.
package auth

import (
	"errors"
	"fmt"
)

// ErrUnauthorized, et al. are custom auth-based errors.
var (
	ErrUnauthorized     = errors.New("unauthorized")
	ErrNotFound         = errors.New("entity not found")
	ErrDuplicateAccount = errors.New("duplicate user account")
	ErrSystem           = errors.New("system error")
)

func systemError(err error) error {
	return fmt.Errorf("%w: %w", ErrSystem, err)
}
