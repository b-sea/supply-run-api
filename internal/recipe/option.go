package recipe

import (
	"errors"
	"net/url"
	"strings"

	"github.com/b-sea/supply-run-api/internal/entity"
)

// Option is a Recipe create or update option.
type Option func(r *Recipe) (bool, error)

// SetName sets the Recipe name.
// Error cases:
//   - Name is empty
func SetName(name string) Option {
	return func(r *Recipe) (bool, error) {
		if name == "" {
			return false, errors.New("recipe name cannot be empty")
		}

		if r.name == name {
			return false, nil
		}

		r.name = name

		return true, nil
	}
}

// SetURL sets the Recipe url.
// Error cases:
//   - URL is not a valid url format
func SetURL(value string) Option {
	return func(r *Recipe) (bool, error) {
		if _, err := url.ParseRequestURI(value); err != nil {
			return false, errors.New("recipe url must be a valid url")
		}

		if r.url == value {
			return false, nil
		}

		r.url = value

		return true, nil
	}
}

// SetNumServings sets the number of servings in a Recipe.
// If the amount is 0 or less, it will be set to 1.
func SetNumServings(num int) Option {
	return func(r *Recipe) (bool, error) {
		if num <= 0 {
			num = 1
		}

		if r.numServings == num {
			return false, nil
		}

		r.numServings = num

		return true, nil
	}
}

// AddStep adds a Recipe step.
func AddStep(step string) Option {
	return func(r *Recipe) (bool, error) {
		r.steps = append(r.steps, step)

		return true, nil
	}
}

// ClearSteps removes all Recipe steps.
func ClearSteps() Option {
	return func(r *Recipe) (bool, error) {
		if len(r.steps) == 0 {
			return false, nil
		}

		r.steps = make([]string, 0)

		return true, nil
	}
}

// AddIngredient adds an ingredient to a Recipe.
// Error cases:
//   - Name is empty
//   - Quantity is 0 or less
func AddIngredient(name string, quantity float64, unitID entity.ID) Option {
	return func(r *Recipe) (bool, error) {
		if name == "" {
			return false, errors.New("ingredient name cannot be empty")
		}

		if quantity <= 0 {
			return false, errors.New("ingredient quantity must be greater than 0")
		}

		r.ingredients = append(
			r.ingredients,
			Ingredient{
				name:     name,
				quantity: quantity,
				unitID:   unitID,
			},
		)

		return true, nil
	}
}

// ClearIngredients removes all Recipe ingredients.
func ClearIngredients() Option {
	return func(r *Recipe) (bool, error) {
		if len(r.ingredients) == 0 {
			return false, nil
		}

		r.ingredients = make([]Ingredient, 0)

		return true, nil
	}
}

// AddTag adds a Recipe tag.
// Error cases:
//   - Name is empty
func AddTag(name string) Option {
	return func(r *Recipe) (bool, error) {
		if name == "" {
			return false, errors.New("tag name cannot be empty")
		}

		for i := range r.tags {
			if !strings.EqualFold(r.tags[i], name) {
				continue
			}

			return false, nil
		}

		r.tags = append(r.tags, name)

		return true, nil
	}
}

// ClearTags removes all Recipe tags.
func ClearTags() Option {
	return func(r *Recipe) (bool, error) {
		if len(r.tags) == 0 {
			return false, nil
		}

		r.tags = make([]string, 0)

		return true, nil
	}
}
