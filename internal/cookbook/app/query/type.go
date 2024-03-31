package query

import (
	"time"

	"github.com/b-sea/supply-run-api/internal/domain/unit"
	"github.com/google/uuid"
)

type RecipeSnippet struct {
	ID          uuid.UUID   `json:"id"`
	CreatedAt   time.Time   `json:"created_at"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	TagIDs      []uuid.UUID `json:"tag_ids"`
}

type RecipeFilter struct {
	Name   *string      `json:"name"`
	TagIDs []uuid.UUIDs `json:"tag_ids"`
}

type Recipe struct {
	ID            uuid.UUID   `json:"id"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     *time.Time  `json:"updated_at"`
	OwnerID       uuid.UUID   `json:"owner_id"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	URL           *string     `json:"url"`
	IngredientIDs []uuid.UUID `json:"ingredient_ids"`
	Servings      int         `json:"servings"`
	Steps         []string    `json:"steps"`
	TagIDs        []uuid.UUID `json:"tag_ids"`
}

type Ingredient struct {
	ID       uuid.UUID `json:"id"`
	ItemID   uuid.UUID `json:"item_id"`
	UnitID   uuid.UUID `json:"unit_id"`
	Quantity float64   `json:"quantity"`
}

type Unit struct {
	ID        uuid.UUID   `json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at"`
	OwnerID   uuid.UUID   `json:"owner_id"`
	Name      string      `json:"name"`
	Symbol    string      `json:"symbol"`
	Type      unit.Type   `json:"type"`
	System    unit.System `json:"system"`
}
