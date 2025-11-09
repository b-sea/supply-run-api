package recipe_test

import (
	"testing"
	"time"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/recipe"
	units "github.com/bcicen/go-units"
	"github.com/stretchr/testify/assert"
)

func TestSetName(t *testing.T) {
	t.Parallel()

	test, err := recipe.New(entity.NewID(), "test", time.Now(), entity.NewID())
	assert.NoError(t, err)

	// Set an empty name
	changed, err := recipe.SetName("")(test)
	assert.False(t, changed)
	assert.Error(t, err)
	assert.Equal(t, "test", test.Name())

	// Set the name
	changed, err = recipe.SetName("new name")(test)
	assert.True(t, changed)
	assert.NoError(t, err)
	assert.Equal(t, "new name", test.Name())

	// Set the name to the same value
	changed, err = recipe.SetName("new name")(test)
	assert.False(t, changed)
	assert.NoError(t, err)
	assert.Equal(t, "new name", test.Name())
}

func TestSetURL(t *testing.T) {
	t.Parallel()

	test, err := recipe.New(entity.NewID(), "test", time.Now(), entity.NewID())
	assert.NoError(t, err)

	// Set the URL
	changed, err := recipe.SetURL("http://test.org/my/recipe")(test)
	assert.True(t, changed)
	assert.NoError(t, err)
	assert.Equal(t, "http://test.org/my/recipe", test.URL())

	// Set a non-URL
	changed, err = recipe.SetURL("i am not.a-url")(test)
	assert.False(t, changed)
	assert.Error(t, err)
	assert.Equal(t, "http://test.org/my/recipe", test.URL())

	// Set the URL to the same value
	changed, err = recipe.SetURL("http://test.org/my/recipe")(test)
	assert.False(t, changed)
	assert.NoError(t, err)
	assert.Equal(t, "http://test.org/my/recipe", test.URL())
}

func TestSetNumServings(t *testing.T) {
	t.Parallel()

	test, err := recipe.New(entity.NewID(), "test", time.Now(), entity.NewID())
	assert.NoError(t, err)

	// Set number of servings
	changed, err := recipe.SetNumServings(4)(test)
	assert.True(t, changed)
	assert.NoError(t, err)
	assert.Equal(t, 4, test.NumServings())

	// Set number of servings to the same value
	changed, err = recipe.SetNumServings(4)(test)
	assert.False(t, changed)
	assert.NoError(t, err)
	assert.Equal(t, 4, test.NumServings())

	// Set number of servings to an invalid value
	changed, err = recipe.SetNumServings(0)(test)
	assert.True(t, changed)
	assert.NoError(t, err)
	assert.Equal(t, 1, test.NumServings())
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

	// Add a duplicate ingredient with a different unit, should be case-insensitive
	changed, err = recipe.AddIngredient("fLoUR", 250, units.Gram)(test)
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

func TestAddTag(t *testing.T) {
	t.Parallel()

	test, err := recipe.New(entity.NewID(), "test", time.Now(), entity.NewID())
	assert.NoError(t, err)

	// Add a tag
	changed, err := recipe.AddTag("tasty")(test)
	assert.True(t, changed)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(test.Tags()))
	assert.Equal(t, "tasty", test.Tags()[0])

	// Add the same tag again with different casing
	changed, err = recipe.AddTag("tAsTy")(test)
	assert.False(t, changed)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(test.Tags()))
	assert.Equal(t, "tasty", test.Tags()[0])

	// Add an empty tag
	changed, err = recipe.AddTag("")(test)
	assert.False(t, changed)
	assert.Error(t, err)

	// Add a different tag
	changed, err = recipe.AddTag("not tasty")(test)
	assert.True(t, changed)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(test.Tags()))
	assert.Equal(t, "not tasty", test.Tags()[1])
}

func TestClearTags(t *testing.T) {
	t.Parallel()

	test, err := recipe.New(
		entity.NewID(), "test", time.Now(), entity.NewID(),
		recipe.AddTag("tasty"),
		recipe.AddTag("not tasty"),
	)
	assert.NoError(t, err)

	// Clearing tags should report a change
	changed, err := recipe.ClearTags()(test)
	assert.True(t, changed)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(test.Tags()))

	// Clearing tags again should report no change
	changed, err = recipe.ClearTags()(test)
	assert.False(t, changed)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(test.Tags()))
}
