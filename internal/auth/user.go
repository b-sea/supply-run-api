package auth

import "github.com/google/uuid"

// User is data about an authenticated user.
type User struct {
	ID       uuid.UUID
	Username string
}
