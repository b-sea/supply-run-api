package query

import (
	"time"

	"github.com/b-sea/supply-run-api/internal/entity"
)

// Recipe is a query representation of a domain Recipe.
type Recipe struct {
	ID          entity.ID
	Name        string
	URL         string
	NumServings int
	Steps       []string
	Ingredients []Ingredient
	Tags        []string
	IsFavorite  bool
	CreatedAt   time.Time
	CreatedBy   entity.ID
	UpdatedAt   time.Time
	UpdatedBy   entity.ID
}

// RecipeFilter defines all options available for finding recipes.
type RecipeFilter struct {
	Name        *string
	Ingredients []string
	CreatedBy   *entity.ID
}

// Ingredient is a query representation of a domain Ingredient.
type Ingredient struct {
}

// User is a query representation of a domain User.
type User struct {
	ID       entity.ID
	Username string
}

// Cursor points to a specific item on a paged result.
type Cursor struct {
	ID  entity.ID
	Key any
}

// Pagination defines pagination controls for queries.
type Pagination struct {
	First  int
	After  Cursor
	Last   int
	Before Cursor
}
