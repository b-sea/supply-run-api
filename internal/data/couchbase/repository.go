package couchbase

import (
	"fmt"
	"strconv"
	"time"

	"github.com/b-sea/supply-run-api/internal/model"
	"github.com/couchbase/gocb/v2"
)

func casToString(cas gocb.Cas) string {
	return strconv.FormatUint(uint64(cas), 10)
}

func stringToCas(cas string) gocb.Cas {
	conv, _ := strconv.ParseUint(cas, 10, 64)
	return gocb.Cas(conv)
}

type dto[E model.Entity] interface {
	Metadata() metadata
	FromEntity(entity E) interface{}
	ToEntity(metadata metadata) E
}

type crudRepository[D dto[E], E model.Entity] struct {
	cluster         *gocb.Cluster
	scope_name      string
	collection_name string
}

func (r *crudRepository[D, E]) GetOne(key string) (E, error) {
	var empty E

	collection := r.cluster.Bucket(bucket_name).Scope(r.scope_name).Collection(r.collection_name)
	doc, err := collection.Get(key, nil)
	if err != nil {
		return empty, fmt.Errorf("error retrieving document: %w", err)
	}

	var dto D
	if err = doc.Content(&dto); err != nil {
		return empty, fmt.Errorf("error parsing document: %w", err)
	}
	metadata := metadata{
		Key:      key,
		Revision: doc.Cas(),
	}
	return dto.ToEntity(metadata), nil
}

func (r *crudRepository[D, E]) upsert(entity E) (string, error) {
	var start D
	dto, ok := start.FromEntity(entity).(D)
	if !ok {
		return "", fmt.Errorf("unable to convert %T to %T", entity, start)
	}

	collection := r.cluster.Bucket(bucket_name).Scope(r.scope_name).Collection(r.collection_name)
	result, err := collection.Upsert(dto.Metadata().Key, dto, nil)
	if err != nil {
		return "", fmt.Errorf("error upserting document: %w", err)
	}

	return casToString(result.Cas()), nil
}

func (r *crudRepository[D, E]) Create(entity E) (string, error) {
	result, err := r.upsert(entity)
	if err != nil {
		return "", fmt.Errorf("error creating document: %w", err)
	}
	return result, nil
}

func (r *crudRepository[D, E]) Update(entity E) (string, error) {
	result, err := r.upsert(entity)
	if err != nil {
		return "", fmt.Errorf("error updating document: %w", err)
	}
	return result, nil
}

func (r *crudRepository[D, E]) Delete(key string) error {
	collection := r.cluster.Bucket(bucket_name).Scope(r.scope_name).Collection(r.collection_name)
	_, err := collection.Remove(key, nil)
	if err != nil {
		return fmt.Errorf("error deleting document: %w", err)
	}
	return nil
}

type metadata struct {
	Key      string   `json:"-"`
	Revision gocb.Cas `json:"-"`
}

type storeDTO struct {
	metadata

	Name string `json:"name"`

	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`

	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`

	Address string `json:"address"`
	Website string `json:"website"`
}

func (d storeDTO) Metadata() metadata {
	return d.metadata
}

func (d storeDTO) FromEntity(entity model.Store) interface{} {
	return storeDTO{
		metadata: metadata{
			Key:      entity.ID.Key,
			Revision: stringToCas(entity.Revision),
		},
		Name: entity.Name,

		CreatedBy: entity.CreatedBy,
		CreatedAt: entity.CreatedAt,

		UpdatedBy: entity.UpdatedBy,
		UpdatedAt: entity.UpdatedAt,

		Address: entity.Address,
		Website: entity.Website,
	}
}

func (d storeDTO) ToEntity(metadata metadata) model.Store {
	return model.Store{
		Metadata: model.Metadata{
			ID: model.GlobalID{
				Key:  metadata.Key,
				Kind: "Store",
			},
			Revision: casToString(metadata.Revision),

			CreatedBy: d.CreatedBy,
			CreatedAt: d.CreatedAt,

			UpdatedBy: d.UpdatedBy,
			UpdatedAt: d.UpdatedAt,
		},
		Name:    d.Name,
		Address: d.Address,
		Website: d.Website,
	}
}

type StoreRepository struct {
	*crudRepository[storeDTO, model.Store]
}

func NewStoreRepository(cluster *gocb.Cluster) *StoreRepository {
	return &StoreRepository{
		crudRepository: &crudRepository[storeDTO, model.Store]{
			cluster:         cluster,
			scope_name:      "entities",
			collection_name: "stores",
		},
	}
}
