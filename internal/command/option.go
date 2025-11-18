package command

import (
	"time"

	"github.com/b-sea/supply-run-api/internal/entity"
)

type Option func(s *Service)

func WithIDGenerator(fn func() entity.ID) Option {
	return func(s *Service) {
		s.idFn = fn
	}
}

func WithTimestampGenerator(fn func() time.Time) Option {
	return func(s *Service) {
		s.timestampFn = fn
	}
}
