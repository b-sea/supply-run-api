package recipe_test

import (
	"testing"
	"time"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/recipe"
	"github.com/stretchr/testify/assert"
)

func TestNewRecipe(t *testing.T) {
	t.Parallel()

	// Create a valid recipe
	id := entity.NewID()
	timestamp := time.Now()
	userID := entity.NewID()
	test, err := recipe.New(id, "test", timestamp, userID)

	assert.NoError(t, err)
	assert.Equal(t, test.ID(), id)
	assert.Equal(t, test.CreatedAt(), timestamp)
	assert.Equal(t, test.CreatedBy(), userID)
	assert.Equal(t, test.UpdatedAt(), timestamp)
	assert.Equal(t, test.UpdatedBy(), userID)

	// Create a recipe with an empty name
	_, err = recipe.New(entity.NewID(), "", time.Now(), entity.NewID())
	assert.Error(t, err)
}

func TestUpdateRecipe(t *testing.T) {
	t.Parallel()

	test, err := recipe.New(entity.NewID(), "test", time.Now(), entity.NewID())
	assert.NoError(t, err)

	// Update the recipe with a valid name
	timestamp := time.Now()
	userID := entity.NewID()
	err = test.Update(timestamp, userID, recipe.SetName("new name"))

	assert.NoError(t, err)
	assert.Equal(t, test.Name(), "new name")
	assert.Equal(t, test.UpdatedAt(), timestamp)
	assert.Equal(t, test.UpdatedBy(), userID)

	// Update the recipe with the same name
	err = test.Update(timestamp, userID, recipe.SetName("new name"))

	assert.NoError(t, err)
	assert.Equal(t, test.Name(), "new name")
	assert.Equal(t, test.UpdatedAt(), timestamp)
	assert.Equal(t, test.UpdatedBy(), userID)

	// Update the recipe with an invalid name
	err = test.Update(time.Now(), entity.NewID(), recipe.SetName(""))

	assert.Error(t, err)
	assert.Equal(t, test.Name(), "new name")
	assert.Equal(t, test.UpdatedAt(), timestamp)
	assert.Equal(t, test.UpdatedBy(), userID)
}
