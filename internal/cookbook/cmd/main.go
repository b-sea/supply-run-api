package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/b-sea/go-auth/token"
	"github.com/b-sea/supply-run-api/config"
	"github.com/b-sea/supply-run-api/internal/auth"
	"github.com/b-sea/supply-run-api/internal/cookbook/adapter"
	"github.com/b-sea/supply-run-api/internal/cookbook/app"
	"github.com/b-sea/supply-run-api/internal/cookbook/app/query"
	"github.com/b-sea/supply-run-api/internal/cookbook/graphql"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func tokenService(cfg *config.Config) *token.Service {
	publicKey, err := os.ReadFile(cfg.JWT.PublicKeyPath)
	if err != nil {
		panic("bad public key")
	}

	privateKey, err := os.ReadFile(cfg.JWT.PrivateKeyPath)
	if err != nil {
		panic("bad private key")
	}

	service, err := token.NewService(
		publicKey,
		privateKey,
		token.WithAccessTimeout(time.Duration(cfg.JWT.AccessTimeout)*time.Hour),
		token.WithRefreshTimeout(time.Duration(cfg.JWT.AccessTimeout)*time.Hour),
		token.WithAudience(cfg.JWT.Audience),
		token.WithIssuer(cfg.JWT.Issuer),
	)
	if err != nil {
		panic("bad token service")
	}

	return service
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	cookbook := adapter.NewCookbookMemoryRepository()
	app := app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			FindRecipeSnippets: query.NewFindRecipeSnippetsHandler(cookbook),
			GetRecipe:          query.NewGetRecipeHandler(cookbook),
			GetIngredients:     query.NewGetIngredientsHandler(cookbook),
			GetAllUnits:        query.NewGetAllUnitsHandler(cookbook),
			GetUnits:           query.NewGetUnitsHandler(cookbook),
		},
	}

	authService := auth.NewService(
		auth.NewMemoryRepository(
			[]*auth.User{
				{ID: uuid.MustParse("51f3027c-5b15-4cf2-a7d2-0c82bb6d7d1a"), Username: "bcarl"},
			},
		),
		tokenService(cfg),
		nil,
	)

	logrus.Warn(authService.Token().GenerateAccessToken("user"))
	logrus.Warn("")
	logrus.Warn(authService.Token().GenerateAccessToken("imposter"))

	schema := graphql.NewExecutableSchema(
		graphql.Config{
			Resolvers: &graphql.Resolver{
				Auth: authService,
				App:  &app,
			},
		},
	)

	srv := handler.NewDefaultServer(schema)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", authService.Middleware(srv))

	port := "8080"
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
