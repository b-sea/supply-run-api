// Main package is the entrypoint for the program
package main

import (
	config "github.com/b-sea/supply-run-api/configs"
	"github.com/b-sea/supply-run-api/internal/data/couchbase"
	"github.com/b-sea/supply-run-api/internal/model"
	"github.com/b-sea/supply-run-api/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := couchbase.NewConnection(*cfg)
	if err != nil {
		logrus.Fatal(err)
	}

	repo := couchbase.NewProductRepository(db)
	svc := service.NewProductService(repo)

	// input := model.CreateProductInput{
	// 	Name:        "bread",
	// 	Description: "it's bread. that's all.",
	// }
	// result, err := svc.Create(input)
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	// logrus.Infof("%+v", result)

	result := &model.CreateResult{
		ID: model.GlobalID{
			Key:  "887941b8-8cc6-4422-b795-1b0f8dd3e8aa",
			Kind: model.ProductKind,
		},
	}
	found, err := svc.GetOne(result.ID)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("%+v", found)

}
