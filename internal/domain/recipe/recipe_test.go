package recipe_test

import (
	"testing"
	"time"

	"github.com/b-sea/supply-run-api/internal/domain/recipe"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRecipeNewRecipe(t *testing.T) {
	t.Parallel()

	id := uuid.New()
	owner := uuid.New()
	now := time.Now()

	tags := []*recipe.Tag{recipe.NewTag("sandwich", recipe.WithTagID(uuid.New()))}
	steps := []*recipe.Step{recipe.NewStep("make sandwich")}
	ingredients := []*recipe.Ingredient{recipe.NewIngredient(uuid.New(), uuid.New(), 1)}

	testRecipe := recipe.NewRecipe(
		"PB&J",
		owner,
		recipe.WithRecipeID(id),
		recipe.WithTimestamp(now),
		recipe.WithDescription("a delicious sandwich"),
		recipe.WithServings(1),
		recipe.WithURL("http://www.pbandj.net"),
		recipe.WithTags(tags),
		recipe.WithSteps(steps),
		recipe.WithIngredients(ingredients),
	)

	assert.Equal(t, id, testRecipe.ID())
	assert.Equal(t, now, testRecipe.CreatedAt())
	assert.Equal(t, owner, testRecipe.Owner())
	assert.Equal(t, "PB&J", testRecipe.Name)
	assert.Equal(t, 1, testRecipe.Servings)
	assert.Equal(t, "http://www.pbandj.net", testRecipe.URL)
	assert.Equal(t, tags, testRecipe.Tags)
	assert.Equal(t, steps, testRecipe.Steps)
	assert.Equal(t, ingredients, testRecipe.Ingredients)
}

func TestRecipeValidate(t *testing.T) {
	t.Parallel()

	testRecipe := recipe.NewRecipe("", uuid.New())
	err := testRecipe.Validate()

	assert.EqualError(t, err, "validation errors: name cannot be empty")

	testRecipe = recipe.NewRecipe("my recipe", uuid.New())
	err = testRecipe.Validate()

	assert.NoError(t, err)
}

func TestRecipeNewStep(t *testing.T) {
	t.Parallel()

	id := uuid.New()
	testStep := recipe.NewStep("make yourself a sandwich", recipe.WithStepID(id))

	assert.Equal(t, id, testStep.ID())
	assert.Equal(t, "make yourself a sandwich", testStep.Details)
}

func TestRecipeStepValidate(t *testing.T) {
	t.Parallel()

	testStep := recipe.NewStep("")
	err := testStep.Validate()

	assert.EqualError(t, err, "validation errors: details cannot be empty")

	testStep = recipe.NewStep("do something")
	err = testStep.Validate()

	assert.NoError(t, err)
}

func TestRecipeNewTag(t *testing.T) {
	t.Parallel()

	id := uuid.New()
	testTag := recipe.NewTag("sandwich", recipe.WithTagID(id))

	assert.Equal(t, id, testTag.ID())
	assert.Equal(t, "sandwich", testTag.Name())
}

func TestRecipeTagValidate(t *testing.T) {
	t.Parallel()

	testTag := recipe.NewTag("")
	err := testTag.Validate()

	assert.EqualError(t, err, "validation errors: name cannot be empty")

	testTag = recipe.NewTag("sandwich")
	err = testTag.Validate()

	assert.NoError(t, err)
}

func TestRecipeNewIngredient(t *testing.T) {
	t.Parallel()

	id := uuid.New()
	item := uuid.New()
	unit := uuid.New()
	testIngredient := recipe.NewIngredient(item, unit, 5, recipe.WithIngredientID(id))

	assert.Equal(t, id, testIngredient.ID())
	assert.Equal(t, item, testIngredient.Item())
	assert.Equal(t, unit, testIngredient.Unit)
	assert.Equal(t, 5.0, testIngredient.Quantity)
}

func TestRecipeIngredientValidate(t *testing.T) {
	t.Parallel()

	testIngredient := recipe.NewIngredient(uuid.New(), uuid.New(), 0)
	err := testIngredient.Validate()

	assert.EqualError(t, err, "validation errors: quantity cannot be 0 or less than 0")

	testIngredient = recipe.NewIngredient(uuid.New(), uuid.New(), 4.654)
	err = testIngredient.Validate()

	assert.NoError(t, err)
}
