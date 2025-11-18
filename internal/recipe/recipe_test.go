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
	id := entity.NewID("recipe-123")
	timestamp := time.Now()
	userID := entity.NewID("user-123")
	test, err := recipe.New(id, "test", timestamp, userID)

	assert.NoError(t, err)
	assert.Equal(t, id, test.ID())
	assert.Equal(t, timestamp.UTC(), test.CreatedAt())
	assert.Equal(t, userID, test.CreatedBy())
	assert.Equal(t, timestamp.UTC(), test.UpdatedAt())
	assert.Equal(t, userID, test.UpdatedBy())

	// Create a recipe with an empty name
	_, err = recipe.New(entity.NewRandomID(), "", time.Now(), entity.NewRandomID())
	assert.Error(t, err)
}

func TestUpdateRecipe(t *testing.T) {
	t.Parallel()

	test, err := recipe.New(entity.NewRandomID(), "test", time.Now(), entity.NewRandomID())
	assert.NoError(t, err)

	// Update the recipe with a valid name
	timestamp := time.Now()
	userID := entity.NewID("user-123")
	err = test.Update(timestamp, userID, recipe.SetName("new name"))

	assert.NoError(t, err)
	assert.Equal(t, "new name", test.Name())
	assert.Equal(t, timestamp.UTC(), test.UpdatedAt())
	assert.Equal(t, userID, test.UpdatedBy())

	// Update the recipe with the same name
	err = test.Update(timestamp, userID, recipe.SetName("new name"))

	assert.NoError(t, err)
	assert.Equal(t, "new name", test.Name())
	assert.Equal(t, timestamp.UTC(), test.UpdatedAt())
	assert.Equal(t, userID, test.UpdatedBy())

	// Update the recipe with an invalid name
	err = test.Update(time.Now(), entity.NewRandomID(), recipe.SetName(""))

	assert.Error(t, err)
	assert.Equal(t, "new name", test.Name())
	assert.Equal(t, timestamp.UTC(), test.UpdatedAt())
	assert.Equal(t, userID, test.UpdatedBy())
}
