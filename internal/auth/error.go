package auth

import (
	"errors"
)

var (
	ErrUnauthorized     = errors.New("unauthorized")
	ErrNotFound         = errors.New("entity not found")
	ErrDuplicateAccount = errors.New("duplicate user account")
)
