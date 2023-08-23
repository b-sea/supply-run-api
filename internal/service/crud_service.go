package service

import (
	"github.com/b-sea/supply-run-api/internal/model"
	"github.com/b-sea/supply-run-api/internal/repository"
)

type IReadService[N model.Node, F any] interface {
	Find(input F) ([]N, error)
	GetOne(id model.GlobalID) (*N, error)
	GetMany(ids []model.GlobalID) ([]*N, error)
}

type IReadWriteService[N model.Node, F any, C model.CreateInput[N], U model.UpdateInput[N]] interface {
	IReadService[N, F]

	Create(input C) (*model.CreateResult, error)
	Update(input U) (*model.UpdateResult, error)
	Delete(input model.DeleteInput) error
}

type crudService[N model.Node, F any, C model.CreateInput[N], U model.UpdateInput[N]] struct {
	repo repository.IReadWriteRepo[N, model.NodeFilter]
	kind model.Kind
}

func (s *crudService[N, F, C, U]) Find(input F) ([]N, error) {
	return nil, nil
}

func (s *crudService[N, F, C, U]) GetOne(id model.GlobalID) (*N, error) {
	if id.Kind != s.kind {
		return nil, model.NewNotFoundError(id)
	}
	result, err := s.repo.GetOne(id.Key)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *crudService[N, F, C, U]) GetMany(ids []model.GlobalID) ([]*N, error) {
	keys := make([]string, len(ids))
	for i, id := range ids {
		if id.Kind != s.kind {
			continue
		}
		keys[i] = id.Key
	}

	result, err := s.repo.GetMany(keys)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *crudService[N, F, C, U]) Create(input C) (*model.CreateResult, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	nodePtr := input.ToNode()

	revision, err := s.repo.Create(nodePtr)
	if err != nil {
		return nil, err
	}

	node := *nodePtr
	result := &model.CreateResult{
		ID:       node.GetID(),
		Revision: revision,
	}
	return result, nil
}

func (s *crudService[N, F, C, U]) Update(input U) (*model.UpdateResult, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	nodePtr, err := s.GetOne(input.GetID())
	if err != nil {
		return nil, err
	}

	node := *nodePtr
	if input.GetRevision() != node.GetRevision() {
		return nil, model.NewConflictError(node.GetRevision())
	}

	input.MergeNode(nodePtr)
	revision, err := s.repo.Update(nodePtr)
	if err != nil {
		return nil, err
	}

	result := &model.UpdateResult{
		ID:       node.GetID(),
		Revision: revision,
	}
	return result, nil
}

func (s *crudService[N, F, C, U]) Delete(input model.DeleteInput) error {
	nodePtr, err := s.GetOne(input.ID)
	if err != nil {
		return err
	}

	node := *nodePtr
	if input.Revision != node.GetRevision() {
		return model.NewConflictError(node.GetRevision())
	}

	err = s.repo.Delete(input.ID.Key)
	if err != nil {
		return err
	}

	return nil
}
