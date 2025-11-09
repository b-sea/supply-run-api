package user_test

import (
	"testing"

	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/user"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	t.Parallel()

	id := entity.NewID()
	test := user.New(id, "tester")

	assert.Equal(t, id, test.ID())
	assert.Equal(t, "tester", test.Username())
}
