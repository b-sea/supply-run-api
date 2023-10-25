// Package local implements memory-based data storage.
package local

import (
	"github.com/b-sea/supply-run-api/internal/model"
	"golang.org/x/exp/slices"
)

type entityRepo[E model.IEntity, F any] struct {
	data        map[string]*E
	filterMatch func(*F, *E) bool
}

func (r *entityRepo[E, F]) Find(filter *F) ([]*E, error) {
	result := []*E{}
	for _, data := range r.data {
		if !r.filterMatch(filter, data) {
			continue
		}
		result = append(result, data)
	}
	return result, nil
}

func (r *entityRepo[E, F]) GetOne(id string) (*E, error) {
	return r.data[id], nil
}

func (r *entityRepo[E, F]) GetMany(ids []string) ([]*E, error) {
	result := []*E{}
	for id, data := range r.data {
		if !slices.Contains(ids, id) {
			continue
		}
		result = append(result, data)
	}
	return result, nil
}

func (r *entityRepo[E, F]) Create(entity E) error {
	r.data[entity.GetID().Key] = &entity
	return nil
}

func (r *entityRepo[E, F]) Update(entity E) error {
	r.data[entity.GetID().Key] = &entity
	return nil
}

func (r *entityRepo[E, F]) Delete(id string) error {
	found, err := r.GetOne(id)
	if err != nil {
		return err
	}

	if found == nil {
		return nil
	}

	delete(r.data, id)
	return nil
}
