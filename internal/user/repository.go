package user

import "github.com/b-sea/supply-run-api/internal/entity"

// Repository defines all data interactions required for users.
type Repository interface {
	GetUser(id entity.ID) (*User, error)
	CreateUser(user *User) error
	DeleteUser(id entity.ID) error
}
