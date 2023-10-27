// Package model defines all data entities shared between front end, service, and repository layers.
package model

import "time"

// RecipeKind is the ID type associated with Accounts.
const RecipeKind = kind("Account")

type Recipe struct {
	ID        ID        `json:"id"`
	AccountID ID        `json:"accountID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Title    string  `json:"title"`
	Servings int     `json:"servings"`
	URL      *string `json:"url"`
	Notes    string  `json:"notes"`
}

// GetID returns the ID of an Account.
func (m Recipe) GetID() ID {
	return m.ID
}

type RecipeFilter struct {
	Title       StringFilter `json:"title"`
	URL         StringFilter `json:"url"`
	Ingredients IDFilter     `json:"ingredients"`
	Tags        IDFilter     `json:"tags"`
}

func (m RecipeFilter) IsFilter() {}

type CreateRecipeInput struct {
	Title    string  `json:"title"`
	Servings int     `json:"servings"`
	URL      *string `json:"url"`
	Notes    string  `json:"notes"`
}

// ToNode converts a CreateAccountInput to an Account node.
func (m CreateRecipeInput) ToNode(key string, timestamp time.Time) Recipe {
	return Recipe{
		ID: ID{
			Kind: AccountKind,
			Key:  key,
		},
		CreatedAt: timestamp,
		UpdatedAt: timestamp,

		Title:    m.Title,
		Servings: m.Servings,
		URL:      m.URL,
		Notes:    m.Notes,
	}
}

type UpdateRecipeInput struct {
	ID       ID      `json:"id"`
	Title    *string `json:"title"`
	Servings *int    `json:"servings"`
	URL      *string `json:"url"`
	Notes    *string `json:"notes"`
}

// GetID returns the ID of an Account.
func (m UpdateRecipeInput) GetID() ID {
	return m.ID
}

// MergeNode applies all UpdateAccountInput values to a Recipe node.
func (m UpdateRecipeInput) MergeNode(node *Recipe, timestamp time.Time) {
	if m.Title != nil {
		node.Title = *m.Title
	}

	if m.Servings != nil {
		node.Servings = *m.Servings
	}

	if m.URL != nil {
		if *m.URL == "" {
			m.URL = nil
		}
		node.URL = m.URL
	}

	if m.Notes != nil {
		node.Notes = *m.Notes
	}

	node.UpdatedAt = timestamp
}
