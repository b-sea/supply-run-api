// Package model defines graphql types.
package model

import (
	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/query"
)

const defaultPageSize = 50

// NewRecipeID creates a graphql Recipe ID.
func NewRecipeID(id entity.ID) ID {
	return ID{
		Key:  id,
		Kind: RecipeKind,
	}
}

// NewRecipe creates a new graphql Recipe.
func NewRecipe(recipe *query.Recipe) *Recipe {
	ingredients := make([]*Ingredient, len(recipe.Ingredients))
	for i := range recipe.Ingredients {
		ingredients[i] = &Ingredient{
			Name: recipe.Ingredients[i].Name,
		}
	}

	return &Recipe{
		ID:          NewRecipeID(recipe.ID),
		Name:        recipe.Name,
		URL:         recipe.URL,
		NumServings: recipe.NumServings,
		Steps:       recipe.Steps,
		Ingredients: ingredients,
		Tags:        recipe.Tags,
		IsFavorite:  recipe.IsFavorite,
		CreatedAt:   recipe.CreatedAt,
		CreatedBy:   NewUserID(recipe.CreatedBy),
		UpdatedAt:   recipe.UpdatedAt,
		UpdatedBy:   NewUserID(recipe.UpdatedBy),
	}
}

// NewQueryRecipeFilter creates a new query RecipeFilter.
func NewQueryRecipeFilter(filter *RecipeFilter) query.RecipeFilter {
	result := query.RecipeFilter{}

	if filter == nil {
		return result
	}

	result.Name = filter.Name
	result.Ingredients = filter.Ingredients

	if filter.CreatedBy != nil {
		result.CreatedBy = &filter.CreatedBy.Key
	}

	result.IsFavorite = filter.IsFavorite

	return result
}

// NewRecipeConnection creates a new graphql RecipeConnection.
func NewRecipeConnection(page *query.RecipePage) *RecipeConnection {
	result := &RecipeConnection{
		PageInfo: &PageInfo{},
		Edges:    make([]*RecipeEdge, 0),
	}

	if page == nil {
		return result
	}

	result.PageInfo.HasNextPage = page.Info.HasNextPage
	result.PageInfo.HasPreviousPage = page.Info.HasPreviousPage

	var sort Sort

	if page.Info.StartCursor != nil {
		sort = newSort(page.Info.StartCursor.Sort)

		result.PageInfo.StartCursor = &Cursor{
			ID:   page.Info.StartCursor.ID,
			Sort: sort,
		}
	}

	if page.Info.EndCursor != nil {
		sort = newSort(page.Info.EndCursor.Sort)

		result.PageInfo.EndCursor = &Cursor{
			ID:   page.Info.EndCursor.ID,
			Sort: sort,
		}
	}

	for _, recipe := range page.Items {
		result.Edges = append(
			result.Edges,
			&RecipeEdge{
				Cursor: Cursor{
					ID:   recipe.ID,
					Sort: sort,
				},
				Node: NewRecipe(recipe),
			},
		)
	}

	return result
}

// NewUserID creates a new graphql User ID.
func NewUserID(id entity.ID) ID {
	return ID{
		Key:  id,
		Kind: UserKind,
	}
}

// NewUser creates a new graphql User.
func NewUser(user *query.User) *User {
	return &User{
		ID:       NewUserID(user.ID),
		Username: user.Username,
	}
}

// NewQueryPagination creates a new query Pagination.
func NewQueryPagination(page *Page) query.Pagination {
	result := query.Pagination{
		Size:   defaultPageSize,
		Cursor: nil,
	}

	if page == nil {
		return result
	}

	if page.First != nil {
		result.Size = *page.First

		if page.After != nil {
			result.Cursor = &query.Cursor{
				ID:   page.After.ID,
				Sort: newQuerySort(page.After.Sort),
			}
		}
	}

	return result
}

// NewQueryOrder creates a new query Order.
func NewQueryOrder(order *Order) query.Order {
	result := query.Order{}

	if order == nil {
		return result
	}

	if order.Sort != nil {
		result.Sort = newQuerySort(*order.Sort)
	}

	if order.Direction != nil {
		result.Direction = newQueryDirection(*order.Direction)
	}

	return result
}

func newSort(sort query.Sort) Sort {
	switch sort {
	case query.NameSort:
		return SortName
	case query.UpdatedSort:
		return SortUpdated
	case query.CreatedSort:
		fallthrough
	default:
		return SortCreated
	}
}

func newQuerySort(sort Sort) query.Sort {
	switch sort {
	case SortName:
		return query.NameSort
	case SortUpdated:
		return query.UpdatedSort
	case SortCreated:
		fallthrough
	default:
		return query.CreatedSort
	}
}

func newQueryDirection(direciton Direction) query.Direction {
	switch direciton {
	case DirectionAsc:
		return query.AscDirection
	case DirectionDesc:
		fallthrough
	default:
		return query.DescDirection
	}
}
