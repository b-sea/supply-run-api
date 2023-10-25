// Package local implements memory-based data storage.
package local

import "github.com/b-sea/supply-run-api/internal/model"

type AccountRepo struct {
	entityRepo[model.Account, model.AccountFilter]
}

func NewAccountRepo() *AccountRepo {
	return &AccountRepo{
		entityRepo: entityRepo[model.Account, model.AccountFilter]{
			filterMatch: accountFilter,
			data:        make(map[string]*model.Account),
		},
	}
}

func accountFilter(filter *model.AccountFilter, entity *model.Account) bool {
	if filter == nil || filter.Username == nil {
		return true
	}

	switch {
	case filter.Username.Eq != nil:
		return entity.Username == *filter.Username.Eq
	default:
		return false
	}
}
