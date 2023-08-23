package couchbase

import (
	"fmt"
	"strconv"
	"time"

	"github.com/b-sea/supply-run-api/internal/model"
	"github.com/couchbase/gocb/v2"
)

type dtoType string

type metadata struct {
	Key      string   `json:"-"`
	Revision gocb.Cas `json:"-"`
}

type shared struct {
	metadata

	Type dtoType `json:"type"`

	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`

	UpdatedBy string    `json:"updatedBy"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (d shared) GetMetadata() metadata {
	return d.metadata
}

type dto[N model.Node] interface {
	GetMetadata() metadata
	FromNode(node *N) interface{}
	ToNode(metadata metadata) *N
}

type crudRepository[D dto[N], N model.Node] struct {
	cluster         *gocb.Cluster
	scope_name      string
	collection_name string
	dtoType         dtoType
}

func (r *crudRepository[D, N]) GetOne(key string) (*N, error) {
	result, err := r.GetMany([]string{key})
	if err != nil {
		return nil, model.NewServerError(err)
	}
	if len(result) != 1 {
		return nil, model.NewNotFoundError(
			model.GlobalID{
				Key:  key,
				Kind: typeToKind(r.dtoType),
			},
		)
	}
	return result[0], nil
}

func (r *crudRepository[D, N]) GetMany(keys []string) ([]*N, error) {
	collection := r.cluster.Bucket(bucketName).Scope(r.scope_name).Collection(r.collection_name)

	ops := make([]gocb.BulkOp, len(keys))
	for i, k := range keys {
		ops[i] = &gocb.GetOp{ID: k}
	}

	err := collection.Do(ops, nil)
	if err != nil {
		return nil, model.NewServerError(err)
	}

	result := make([]*N, len(keys))
	for i, op := range ops {
		var dto D

		doc := op.(*gocb.GetOp)
		if err = doc.Result.Content(&dto); err != nil {
			return nil, model.NewServerError(err)
		}

		metadata := metadata{
			Key:      doc.ID,
			Revision: doc.Result.Cas(),
		}
		node := dto.ToNode(metadata)
		result[i] = node
	}
	return result, nil
}

func (r *crudRepository[D, N]) upsert(node *N) (string, error) {
	var start D
	dto, ok := start.FromNode(node).(D)
	if !ok {
		return "", model.NewServerError(fmt.Errorf("conversion error %T -> %T", node, start))
	}

	collection := r.cluster.Bucket(bucketName).Scope(r.scope_name).Collection(r.collection_name)
	result, err := collection.Upsert(dto.GetMetadata().Key, dto, nil)
	if err != nil {
		return "", model.NewServerError(err)
	}

	return casToString(result.Cas()), nil
}

func (r *crudRepository[D, N]) Create(node *N) (string, error) {
	result, err := r.upsert(node)
	if err != nil {
		return "", model.NewServerError(err)
	}
	return result, nil
}

func (r *crudRepository[D, N]) Update(node *N) (string, error) {
	result, err := r.upsert(node)
	if err != nil {
		return "", model.NewServerError(err)
	}
	return result, nil
}

func (r *crudRepository[D, N]) Delete(key string) error {
	collection := r.cluster.Bucket(bucketName).Scope(r.scope_name).Collection(r.collection_name)
	_, err := collection.Remove(key, nil)
	if err != nil {
		return model.NewServerError(err)
	}
	return nil
}

func casToString(cas gocb.Cas) string {
	return strconv.FormatUint(uint64(cas), 10)
}

func stringToCas(cas string) gocb.Cas {
	conv, _ := strconv.ParseUint(cas, 10, 64)
	return gocb.Cas(conv)
}

func typeToKind(dt dtoType) model.Kind {
	switch {
	case dt == brandType:
		return model.BrandKind
	case dt == brandProductType:
		return model.BrandProductKind
	case dt == categoryType:
		return model.CategoryKind
	case dt == listItemType:
		return model.ListItemKind
	case dt == productType:
		return model.ProductKind
	case dt == shoppingListType:
		return model.ShoppingListKind
	default:
		return model.Kind(dt)
	}
}

func kindToType(k model.Kind) dtoType {
	switch {
	case k == model.BrandKind:
		return brandType
	case k == model.BrandProductKind:
		return brandProductType
	case k == model.CategoryKind:
		return categoryType
	case k == model.ListItemKind:
		return listItemType
	case k == model.ProductKind:
		return productType
	case k == model.ShoppingListKind:
		return shoppingListType
	default:
		return dtoType(k)
	}
}
