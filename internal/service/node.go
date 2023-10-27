// Package service implements all business logic for the API.
package service

import (
	"context"
	"fmt"
	"time"

	"github.com/b-sea/supply-run-api/internal/model"
	"github.com/b-sea/supply-run-api/internal/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	idGenerator = uuid.NewString
	timestamp   = time.Now().UTC
)

type INode[N model.Node, F model.Filter, C model.CreateInput[N], U model.UpdateInput[N]] interface {
	Find(ctx context.Context, filter *F) ([]*N, error)
	GetByID(ctx context.Context, id model.ID) (*N, error)
	GetByIDs(ctx context.Context, id []model.ID) ([]*N, error)
	Create(ctx context.Context, input C) error
	Update(ctx context.Context, input U) error
	Delete(ctx context.Context, id model.ID) error
}

type node[N model.Node, F model.Filter, C model.CreateInput[N], U model.UpdateInput[N]] struct {
	repo repository.INode[N, F]

	idGenerator func() string
	timestamp   func() time.Time

	createModifier func(node *N) error
	updateModifier func(node *N) error
}

func (s *node[N, F, C, U]) Find(ctx context.Context, filter *F) ([]*N, error) {
	accountID, ok := AccountIDFromContext(ctx)
	if !ok {
		return nil, ErrAuthentication
	}

	nodes, err := s.repo.Find(filter, accountID.Key)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("%w", err)
	}

	return nodes, nil
}

func (s *node[N, F, C, U]) GetByID(ctx context.Context, id model.ID) (*N, error) {
	nodes, err := s.GetByIDs(ctx, []model.ID{id})
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("%w", err)
	}

	if len(nodes) != 1 {
		return nil, model.NotFoundError{ID: id}
	}

	return nodes[0], nil
}

func (s *node[N, F, C, U]) GetByIDs(ctx context.Context, ids []model.ID) ([]*N, error) {
	accountID, ok := AccountIDFromContext(ctx)
	if !ok {
		return nil, ErrAuthentication
	}

	keys := make([]string, len(ids))
	for i := range ids {
		keys[i] = ids[i].Key
	}

	nodes, err := s.repo.GetByIDs(keys, accountID.Key)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("%w", err)
	}

	return nodes, nil
}

func (s *node[N, F, C, U]) Create(ctx context.Context, input C) error {
	accountID, ok := AccountIDFromContext(ctx)
	if !ok {
		return ErrAuthentication
	}

	node := input.ToNode(s.idGenerator(), s.timestamp())

	if s.createModifier != nil {
		if err := s.createModifier(&node); err != nil {
			logrus.Error(err)
			return fmt.Errorf("%w", err)
		}
	}

	if err := s.repo.Create(node, accountID.Key); err != nil {
		logrus.Error(err)
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (s *node[N, F, C, U]) Update(ctx context.Context, input U) error {
	node, err := s.GetByID(ctx, input.GetID())
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	input.MergeNode(node, s.timestamp())

	if s.updateModifier != nil {
		if err := s.updateModifier(node); err != nil {
			logrus.Error(err)
			return fmt.Errorf("%w", err)
		}
	}

	accountID, _ := AccountIDFromContext(ctx)
	if err := s.repo.Update(*node, accountID.Key); err != nil {
		logrus.Error(err)
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (s *node[N, F, C, U]) Delete(ctx context.Context, id model.ID) error {
	accountID, ok := AccountIDFromContext(ctx)
	if !ok {
		return ErrAuthentication
	}

	if err := s.repo.Delete(id.Key, accountID.Key); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
