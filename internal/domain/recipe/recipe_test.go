package recipe_test

import (
	"testing"
)

func TestRecipeNewRecipe(t *testing.T) {
	t.Parallel()

	// id := uuid.New()
	// owner := uuid.New()
	// now := time.Now()

	// tags := []*recipe.Tag{recipe.NewTag("sandwich", recipe.WithTagID(uuid.New()))}
	// steps := []*recipe.Step{recipe.NewStep("make sandwich")}
	// ingredients := []*recipe.Ingredient{recipe.NewIngredient(uuid.New(), uuid.New(), 1)}

	// testRecipe, _ := recipe.NewRecipe(
	// 	"PB&J",
	// 	owner,
	// 	recipe.WithRecipeID(id),
	// 	recipe.WithTimestamp(now),
	// 	recipe.SetDescription("a delicious sandwich"),
	// 	recipe.SetServings(1),
	// 	recipe.SetURL("http://www.pbandj.net"),
	// 	recipe.WithTags(tags),
	// 	recipe.WithSteps(steps),
	// 	recipe.WithIngredients(ingredients),
	// )

	// logrus.Info(testRecipe)
	// err := testRecipe.Update(
	// 	recipe.SetDescription("a delicious sandwich"),
	// 	recipe.SetServings(1),
	// 	recipe.SetURL("http://www.pbandj.net"),
	// 	recipe.SetServings(0),
	// )

	// logrus.Info(">>>>", err)
	// tag, _ := recipe.NewTag("sandwich", owner)
	// logrus.Info(testRecipe.Tags())
	// testRecipe.Update(recipe.SetTags([]*recipe.Tag{tag}))
	// logrus.Info(testRecipe.Tags())
	// logrus.Infof("%+v", testRecipe)

	// assert.Equal(t, id, testRecipe.ID())
	// assert.Equal(t, now, testRecipe.CreatedAt())
	// assert.Equal(t, owner, testRecipe.Owner())
	// assert.Equal(t, "PB&J", testRecipe.Name)
	// assert.Equal(t, 1, testRecipe.Servings)
	// assert.Equal(t, "http://www.pbandj.net", testRecipe.URL)
	// assert.Equal(t, tags, testRecipe.Tags)
	// assert.Equal(t, steps, testRecipe.Steps)
	// assert.Equal(t, ingredients, testRecipe.Ingredients)
}

// func TestRecipeValidate(t *testing.T) {
// 	t.Parallel()

// 	testRecipe := recipe.NewRecipe("", uuid.New())
// 	err := testRecipe.Validate()

// 	assert.EqualError(t, err, "validation errors: name cannot be empty")

// 	testRecipe = recipe.NewRecipe("my recipe", uuid.New())
// 	err = testRecipe.Validate()

// 	assert.NoError(t, err)
// }

// func TestRecipeNewStep(t *testing.T) {
// 	t.Parallel()

// 	id := uuid.New()
// 	testStep := recipe.NewStep("make yourself a sandwich", recipe.WithStepID(id))

// 	assert.Equal(t, id, testStep.ID())
// 	assert.Equal(t, "make yourself a sandwich", testStep.Details)
// }

// func TestRecipeStepValidate(t *testing.T) {
// 	t.Parallel()

// 	testStep := recipe.NewStep("")
// 	err := testStep.Validate()

// 	assert.EqualError(t, err, "validation errors: details cannot be empty")

// 	testStep = recipe.NewStep("do something")
// 	err = testStep.Validate()

// 	assert.NoError(t, err)
// }

// func TestRecipeNewTag(t *testing.T) {
// 	t.Parallel()

// 	id := uuid.New()
// 	testTag := recipe.NewTag("sandwich", recipe.WithTagID(id))

// 	assert.Equal(t, id, testTag.ID())
// 	assert.Equal(t, "sandwich", testTag.Name())
// }

// func TestRecipeTagValidate(t *testing.T) {
// 	t.Parallel()

// 	testTag := recipe.NewTag("")
// 	err := testTag.Validate()

// 	assert.EqualError(t, err, "validation errors: name cannot be empty")

// 	testTag = recipe.NewTag("sandwich")
// 	err = testTag.Validate()

// 	assert.NoError(t, err)
// }

// func TestRecipeNewIngredient(t *testing.T) {
// 	t.Parallel()

// 	id := uuid.New()
// 	item := uuid.New()
// 	unit := uuid.New()
// 	testIngredient := recipe.NewIngredient(item, unit, 5, recipe.WithIngredientID(id))

// 	assert.Equal(t, id, testIngredient.ID())
// 	assert.Equal(t, item, testIngredient.Item())
// 	assert.Equal(t, unit, testIngredient.Unit)
// 	assert.Equal(t, 5.0, testIngredient.Quantity)
// }

// func TestRecipeIngredientValidate(t *testing.T) {
// 	t.Parallel()

// 	testIngredient := recipe.NewIngredient(uuid.New(), uuid.New(), 0)
// 	err := testIngredient.Validate()

// 	assert.EqualError(t, err, "validation errors: quantity cannot be 0 or less than 0")

// 	testIngredient = recipe.NewIngredient(uuid.New(), uuid.New(), 4.654)
// 	err = testIngredient.Validate()

// 	assert.NoError(t, err)
// }
