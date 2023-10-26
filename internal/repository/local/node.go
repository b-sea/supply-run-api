// Package local implements memory-based data storage.
package local

import (
	"github.com/b-sea/supply-run-api/internal/model"
	"golang.org/x/exp/slices"
)

type nodeRepo[N model.Node, F any] struct {
	data        map[string]*N
	filterMatch func(*F, *N) bool
}

func (r *nodeRepo[N, F]) Find(filter *F) ([]*N, error) {
	result := []*N{}
	for _, data := range r.data {
		if !r.filterMatch(filter, data) {
			continue
		}
		result = append(result, data)
	}
	return result, nil
}

func (r *nodeRepo[N, F]) GetOne(id string) (*N, error) {
	return r.data[id], nil
}

func (r *nodeRepo[N, F]) GetMany(ids []string) ([]*N, error) {
	result := []*N{}
	for id, data := range r.data {
		if !slices.Contains(ids, id) {
			continue
		}
		result = append(result, data)
	}
	return result, nil
}

func (r *nodeRepo[N, F]) Create(node N) error {
	r.data[node.GetID().Key] = &node
	return nil
}

func (r *nodeRepo[N, F]) Update(node N) error {
	r.data[node.GetID().Key] = &node
	return nil
}

func (r *nodeRepo[N, F]) Delete(id string) error {
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
