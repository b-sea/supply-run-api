package recipe

import (
	"errors"
	"fmt"

	units "github.com/bcicen/go-units"
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
// If the ingredient already exists on the Recipe, they will be combined.
// Error cases:
//   - Ingredient units cannot be converted when combining
//   - Name is empty
//   - Quantity is 0 or less
func AddIngredient(name string, quantity float64, unit units.Unit) Option {
	return func(r *Recipe) (bool, error) {
		if name == "" {
			return false, errors.New("ingredient name cannot be empty")
		}

		if quantity <= 0 {
			return false, errors.New("ingredient quantity must be greater than 0")
		}

		value := units.NewValue(quantity, unit)
		appended := false

		for i := range r.ingredients {
			if r.ingredients[i].name != name {
				continue
			}

			converted, err := value.Convert(r.ingredients[i].value.Unit())
			if err != nil {
				return false, fmt.Errorf(
					"cannot add %s to %s of %s",
					value.String(),
					r.ingredients[i].value.String(),
					name,
				)
			}

			r.ingredients[i].value = units.NewValue(
				r.ingredients[i].value.Float()+converted.Float(),
				r.ingredients[i].value.Unit(),
			)
			appended = true
		}

		if !appended {
			r.ingredients = append(r.ingredients, Ingredient{name: name, value: value})
		}

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
