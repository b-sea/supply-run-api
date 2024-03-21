package repository

import (
	"github.com/b-sea/supply-run-api/internal/domain/recipe"
	"github.com/b-sea/supply-run-api/internal/domain/tag"
	"github.com/b-sea/supply-run-api/internal/domain/unit"
	"github.com/google/uuid"
)

type RecipeRead interface {
	Find(name *string, tagIDs []uuid.UUID) ([]*recipe.Snippet, error)
	GetByID(id uuid.UUID) (*recipe.Recipe, error)
}

type RecipeWrite interface {
	Create(entity *recipe.Recipe) error
	Update(entity *recipe.Recipe) error
	Delete(id uuid.UUID) error
}

type TagRead interface {
	GetAll() ([]*tag.Tag, error)
}

type TagWrite interface {
	Delete(id uuid.UUID) error
}

type UnitRead interface {
	GetAll() ([]*unit.Unit, error)
}
