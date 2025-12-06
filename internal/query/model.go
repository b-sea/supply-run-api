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

// Ingredient is a query representation of a domain Ingredient.
type Ingredient struct {
	Name     string
	Quantity float64
	UnitID   entity.ID
}

// RecipeFilter defines all options available for finding recipes.
type RecipeFilter struct {
	Name        *string
	Ingredients []string
	CreatedBy   *entity.ID
	IsFavorite  *bool
}

// RecipePage contains information about a page of recipes.
type RecipePage struct {
	Info  PageInfo
	Items []*Recipe
}

// User is a query representation of a domain User.
type User struct {
	ID       entity.ID
	Username string
}

// Unit is a query representation of a domain Unit.
type Unit struct {
	ID       entity.ID
	Name     string
	Symbol   string
	BaseType string
	System   string
}

// Conversion is a query representation of a domain Conversion.
type Conversion struct {
	FromID entity.ID
	ToID   entity.ID
	Ratio  float64
}

// Direction is a sort direction.
type Direction int

// AscDirection, et al. are the different sort directions.
const (
	DescDirection Direction = iota
	AscDirection
)

// Sort is the attribute to sort on.
type Sort int

// NameSort, et al. are the various attributes available for sorting.
const (
	CreatedSort Sort = iota
	UpdatedSort
	NameSort
)

// Order defines ordering information.
type Order struct {
	Sort      Sort
	Direction Direction
}

// Cursor points to a specific item on a paged result.
type Cursor struct {
	ID   entity.ID
	Sort Sort
}

// Pagination defines pagination controls for queries.
type Pagination struct {
	Cursor *Cursor
	Size   int
}

// PageInfo defines the boundries of a page.
type PageInfo struct {
	HasNextPage     bool
	HasPreviousPage bool
	StartCursor     *Cursor
	EndCursor       *Cursor
}
