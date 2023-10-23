package model

import (
	"time"
)

type Account struct {
	ID        ID        `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Revision  string    `json:"revision"`

	Username string `json:"username"`
	Password string `json:"-"`
}

type AccountFilter struct {
	Username StringFilter `json:"username"`
}

type CreateAccountInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResult struct {
	ID       ID     `json:"id"`
	Revision string `json:"revision"`
}

type UpdateAccountInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateResult struct {
	Revision string `json:"revision"`
}

type TokenSet struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}
