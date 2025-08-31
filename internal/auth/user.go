package auth

import "github.com/b-sea/supply-run-api/internal/entity"

type Tokens struct {
	Access  string
	Refresh string
}

type User struct {
	id       entity.ID
	username string
}

func NewUser(id entity.ID, username string) *User {
	return &User{
		id:       id,
		username: username,
	}
}

func (u *User) ID() entity.ID {
	return u.id
}

func (u *User) Username() string {
	return u.username
}
