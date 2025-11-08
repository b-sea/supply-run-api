package recipe_test

import (
	"testing"
	"time"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/recipe"
	units "github.com/bcicen/go-units"
	"github.com/stretchr/testify/assert"
)

func TestRecipeNewRecipe(t *testing.T) {
	t.Parallel()

	type testCase struct {
		err error
	}

	testCases := map[string]testCase{
		"success": {
			err: nil,
		},
	}

	for name, _ := range testCases {
		t.Run(name, func(t *testing.T) {
			// _, err = csplatform.NewRepository("", connector, test.options...)

			// if test.err == nil {
			// 	assert.NoError(t, err)
			// } else {
			// 	assert.ErrorIs(t, err, test.err)
			// }
		})
	}
}

func TestSetName(t *testing.T) {
	t.Parallel()

	test, err := recipe.New(entity.NewID(), "test", time.Now(), entity.NewID())
	assert.NoError(t, err)

	// Set an empty name
	changed, err := recipe.SetName("")(test)
	assert.False(t, changed)
	assert.Error(t, err)
	assert.Equal(t, test.Name(), "test")

	// Set the name
	changed, err = recipe.SetName("new name")(test)
	assert.True(t, changed)
	assert.NoError(t, err)
	assert.Equal(t, test.Name(), "new name")

	// Set the name to the same value
	changed, err = recipe.SetName("new name")(test)
	assert.False(t, changed)
	assert.NoError(t, err)
	assert.Equal(t, test.Name(), "new name")
}

func TestAddStep(t *testing.T) {
	t.Parallel()

	test, err := recipe.New(entity.NewID(), "test", time.Now(), entity.NewID())
	assert.NoError(t, err)

	// Add a step
	changed, err := recipe.AddStep("make sandwich")(test)
	assert.True(t, changed)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(test.Steps()))
	assert.Equal(t, "make sandwich", test.Steps()[0])

	// Add another step
	changed, err = recipe.AddStep("eat sandwich")(test)
	assert.True(t, changed)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(test.Steps()))
	assert.Equal(t, "eat sandwich", test.Steps()[1])
}

func TestClearSteps(t *testing.T) {
	t.Parallel()

	test, err := recipe.New(
		entity.NewID(), "test", time.Now(), entity.NewID(),
		recipe.AddStep("make sandwich"),
		recipe.AddStep("eat sandwich"),
	)
	assert.NoError(t, err)

	// Clearing steps should report a change
	changed, err := recipe.ClearSteps()(test)
	assert.True(t, changed)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(test.Steps()))

	// Clearing steps again should report no change
	changed, err = recipe.ClearSteps()(test)
	assert.False(t, changed)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(test.Steps()))
}

func TestSetNumServings(t *testing.T) {
	t.Parallel()

	test, err := recipe.New(entity.NewID(), "test", time.Now(), entity.NewID())
	assert.NoError(t, err)

	// Set number of servings
	changed, err := recipe.SetNumServings(4)(test)
	assert.True(t, changed)
	assert.NoError(t, err)
	assert.Equal(t, test.NumServings(), 4)

	// Set number of servings to the same value
	changed, err = recipe.SetNumServings(4)(test)
	assert.False(t, changed)
	assert.NoError(t, err)
	assert.Equal(t, test.NumServings(), 4)

	// Set number of servings to an invalid value
	changed, err = recipe.SetNumServings(0)(test)
	assert.True(t, changed)
	assert.NoError(t, err)
	assert.Equal(t, test.NumServings(), 1)
}

func TestAddIngredient(t *testing.T) {
	t.Parallel()

	test, err := recipe.New(entity.NewID(), "test", time.Now(), entity.NewID())
	assert.NoError(t, err)

	// Add an ingredient
	changed, err := recipe.AddIngredient("flour", 2, units.Pound)(test)
	assert.True(t, changed)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(test.Ingredients()))
	assert.Equal(t, "flour", test.Ingredients()[0].Name())
	assert.Equal(t, float64(2), test.Ingredients()[0].Quantity())
	assert.Equal(t, units.Pound, test.Ingredients()[0].Unit())

	// Add a different ingredient
	changed, err = recipe.AddIngredient("water", 6, units.Liter)(test)
	assert.True(t, changed)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(test.Ingredients()))
	assert.Equal(t, "water", test.Ingredients()[1].Name())
	assert.Equal(t, float64(6), test.Ingredients()[1].Quantity())
	assert.Equal(t, units.Liter, test.Ingredients()[1].Unit())

	// Add a duplicate ingredient with a different unit
	changed, err = recipe.AddIngredient("flour", 250, units.Gram)(test)
	assert.True(t, changed)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(test.Ingredients()))
	assert.Equal(t, "flour", test.Ingredients()[0].Name())
	assert.Equal(t, 2.551155655462194, test.Ingredients()[0].Quantity())
	assert.Equal(t, units.Pound, test.Ingredients()[0].Unit())

	// Add a duplicate ingredient with an incompatible unit
	changed, err = recipe.AddIngredient("water", 2, units.Celsius)(test)
	assert.False(t, changed)
	assert.Error(t, err)

	assert.Equal(t, 2, len(test.Ingredients()))

	// Add a ingredient with no name
	changed, err = recipe.AddIngredient("", 1, units.KiloGram)(test)
	assert.False(t, changed)
	assert.Error(t, err)

	assert.Equal(t, 2, len(test.Ingredients()))

	// Add a ingredient with bad quantity
	changed, err = recipe.AddIngredient("something", 0, units.KiloGram)(test)
	assert.False(t, changed)
	assert.Error(t, err)

	assert.Equal(t, 2, len(test.Ingredients()))
}

func TestClearIngredients(t *testing.T) {
	t.Parallel()

	test, err := recipe.New(
		entity.NewID(), "test", time.Now(), entity.NewID(),
		recipe.AddIngredient("flour", 2, units.Pound),
		recipe.AddIngredient("water", 6, units.Liter),
	)
	assert.NoError(t, err)

	// Clearing ingredients should report a change
	changed, err := recipe.ClearIngredients()(test)
	assert.True(t, changed)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(test.Ingredients()))

	// Clearing ingredients again should report no change
	changed, err = recipe.ClearIngredients()(test)
	assert.False(t, changed)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(test.Ingredients()))
}
