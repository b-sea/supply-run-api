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
	IsFavorite  *bool
}

// Ingredient is a query representation of a domain Ingredient.
type Ingredient struct {
	Name string
}

// User is a query representation of a domain User.
type User struct {
	ID       entity.ID
	Username string
}

// Direction is a sort direction.
type Direction int

// AscendingDirection, et al. are the different sort directions.
const (
	AscendingDirection Direction = iota
	DescendingDirection
)

// Sort is the attribute to sort on.
type Sort int

// NameSort, et al. are the various attributes available for sorting.
const (
	NoSort Sort = iota
	NameSort
	CreatedSort
	UpdatedSort
)

func (s Sort) String() string {
	switch s {
	case NameSort:
		return "name"
	case CreatedSort:
		return "created"
	case UpdatedSort:
		return "updated"
	case NoSort:
		fallthrough
	default:
		return ""
	}
}

// Order defines ordering information.
type Order struct {
	Sort      Sort
	Direction Direction
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
