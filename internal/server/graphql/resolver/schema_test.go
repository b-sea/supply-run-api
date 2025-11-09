package resolver_test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/b-sea/supply-run-api/internal/server/graphql/resolver"
	"github.com/stretchr/testify/assert"
)

func TestQueryTemp(t *testing.T) {
	t.Parallel()

	server := handler.New(resolver.NewExecutableSchema(resolver.Config{Resolvers: resolver.NewResolver()}))
	server.AddTransport(transport.POST{})

	testClient := client.New(server)

	var response struct {
		Temp string `json:"temp"`
	}

	assert.NoError(t, testClient.Post(`query { temp }`, &response))
	assert.Equal(t, "query", response.Temp)

}

func TestMutationTemp(t *testing.T) {
	t.Parallel()

	server := handler.New(resolver.NewExecutableSchema(resolver.Config{Resolvers: resolver.NewResolver()}))
	server.AddTransport(transport.POST{})

	testClient := client.New(server)

	var response struct {
		Temp string `json:"temp"`
	}

	assert.NoError(t, testClient.Post(`mutation { temp }`, &response))
	assert.Equal(t, "mutation", response.Temp)

}
