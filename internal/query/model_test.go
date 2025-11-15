package query_test

import (
	"testing"

	"github.com/b-sea/supply-run-api/internal/query"
	"github.com/stretchr/testify/assert"
)

func TestSortString(t *testing.T) {
	t.Parallel()

	type testCase struct {
		sort   query.Sort
		result string
	}

	tests := map[string]testCase{
		"none": {
			sort:   query.NoSort,
			result: "",
		},
		"name": {
			sort:   query.NameSort,
			result: "name",
		},
		"created": {
			sort:   query.CreatedSort,
			result: "created",
		},
		"updated": {
			sort:   query.UpdatedSort,
			result: "updated",
		},
		"unknown": {
			sort:   999,
			result: "",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.result, test.sort.String())
		})
	}
}
