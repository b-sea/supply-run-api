package model_test

import (
	"bytes"
	"testing"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/graphql/model"
	"github.com/stretchr/testify/assert"
)

func TestMarshalID(t *testing.T) {
	t.Parallel()

	result := new(bytes.Buffer)

	marshaler := model.MarshalID(model.NewRecipeID(entity.NewID("recipe-1234")))
	marshaler.MarshalGQL(result)
	assert.Equal(t, `"cmVjaXBlOnJlY2lwZS0xMjM0"`, result.String())
}

func TestUnmarshalID(t *testing.T) {
	t.Parallel()

	type testCase struct {
		value  any
		result model.ID
	}

	tests := map[string]testCase{
		"success": {
			value:  `cmVjaXBlOnJlY2lwZS0xMjM0`,
			result: model.NewRecipeID(entity.NewID("recipe-1234")),
		},
		"bad type": {
			value:  43,
			result: model.ID{},
		},
		"bad encoding": {
			value:  `i am not base 64`,
			result: model.ID{},
		},
		"bad format": {
			value:  `YWJjZGVmZ2hpag==`,
			result: model.ID{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := model.UnmarshalID(test.value)

			assert.Equal(t, test.result, result)
			assert.NoError(t, err)
		})
	}
}

func TestMarshalCursor(t *testing.T) {
	t.Parallel()

	result := new(bytes.Buffer)

	marshaler := model.MarshalCursor(
		model.Cursor{
			ID:   entity.NewID("recipe-1234"),
			Sort: model.SortName,
		},
	)
	marshaler.MarshalGQL(result)
	assert.Equal(t, `"cmVjaXBlLTEyMzQ6TkFNRQ=="`, result.String())
}

func TestUnmarshalCursor(t *testing.T) {
	t.Parallel()

	type testCase struct {
		value  any
		result model.Cursor
	}

	tests := map[string]testCase{
		"success": {
			value: `cmVjaXBlLTEyMzQ6TkFNRQ==`,
			result: model.Cursor{
				ID:   entity.NewID("recipe-1234"),
				Sort: model.SortName,
			},
		},
		"bad type": {
			value:  43,
			result: model.Cursor{},
		},
		"bad encoding": {
			value:  `i am not base 64`,
			result: model.Cursor{},
		},
		"bad format": {
			value:  `YWJjZGVmZ2hpag==`,
			result: model.Cursor{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := model.UnmarshalCursor(test.value)

			assert.Equal(t, test.result, result)
			assert.NoError(t, err)
		})
	}
}
