// Package local implements memory-based data storage.
package local

import "github.com/b-sea/supply-run-api/internal/model"

type AccountRepo struct {
	nodeRepo[model.Account, model.AccountFilter]
}

func NewAccountRepo() *AccountRepo {
	return &AccountRepo{
		nodeRepo: nodeRepo[model.Account, model.AccountFilter]{
			filterMatch: accountFilter,
			data:        make(map[string]*model.Account),
		},
	}
}

func accountFilter(filter *model.AccountFilter, node *model.Account) bool {
	if filter == nil || filter.Email == nil {
		return true
	}

	switch {
	case filter.Email.Eq != nil:
		return node.Email == *filter.Email.Eq
	default:
		return false
	}
}
