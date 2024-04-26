package recipe_test

import (
	"testing"
	"time"

	"github.com/b-sea/supply-run-api/internal/cookbook/domain"
	"github.com/b-sea/supply-run-api/internal/cookbook/domain/recipe"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRecipeNewRecipe(t *testing.T) {
	t.Parallel()

	type test struct {
		id        uuid.UUID
		timestamp time.Time
		ownerID   uuid.UUID
		name      string
		err       error
	}

	testCases := map[string]test{
		"success": {
			name:      "A Fancy Sandwich",
			timestamp: time.Now(),
			ownerID:   uuid.New(),
		},
		"invalid": {
			name: "",
			err:  &domain.ValidationError{Issues: []string{"recipe name cannot be empty"}},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(s *testing.T) {
			s.Parallel()

			newRecipe, err := recipe.NewRecipe(testCase.id, testCase.name, testCase.ownerID, testCase.timestamp)
			if err == nil {
				assert.Equal(t, testCase.id, newRecipe.ID(), "different ID")
				assert.Equal(t, testCase.timestamp.UTC(), newRecipe.CreatedAt(), "different CreatedAt")
				assert.Equal(t, testCase.ownerID, newRecipe.OwnerID(), "different OwnerID")
				assert.Equal(t, testCase.name, newRecipe.Name(), "different Name")
			}

			if testCase.err == nil {
				assert.NoError(t, err, "no error expected")
			} else {
				assert.EqualError(t, err, testCase.err.Error(), "different errors")
			}
		})
	}
}

func TestRecipeUpdate(t *testing.T) {
	t.Parallel()

	newRecipe, err := recipe.NewRecipe(uuid.New(), "A Fancy Sandwich", uuid.New(), time.Now())
	if err != nil {
		panic("creating test recipe failed")
	}

	type test struct {
		timestamp time.Time
		opt       recipe.Option
		err       error
	}

	testCases := map[string]test{
		"SetName success": {
			timestamp: time.Now(),
			opt:       recipe.SetName("PB&J"),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(s *testing.T) {
			s.Parallel()

			err := newRecipe.Update(testCase.timestamp, testCase.opt)
			if err == nil {
				updatedAt := testCase.timestamp.UTC()
				assert.Equal(t, &updatedAt, newRecipe.UpdatedAt(), "different results")
			}

			if testCase.err == nil {
				assert.NoError(t, err, "no error expected")
			} else {
				assert.EqualError(t, err, testCase.err.Error(), "different errors")
			}
		})
	}
}
