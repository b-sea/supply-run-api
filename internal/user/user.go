// Package user defines users.
package user

import "github.com/b-sea/supply-run-api/internal/entity"

// User is a user is a user.
type User struct {
	id       entity.ID
	username string
}

// New creates a new User.
func New(id entity.ID, username string) *User {
	return &User{
		id:       id,
		username: username,
	}
}

// ID gets the User id.
func (u *User) ID() entity.ID {
	return u.id
}

// Username gets the username of the User.
func (u *User) Username() string {
	return u.username
}
