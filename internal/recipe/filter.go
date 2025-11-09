package recipe

import "github.com/b-sea/supply-run-api/internal/entity"

// Filter defines the filterable attributes when finding recipes.
type Filter struct {
	Name        *string
	Ingredients []string
	CreatedBy   *entity.ID
}
