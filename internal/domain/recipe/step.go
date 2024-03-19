// Package recipe defines everything to manage the recipes domain.
package recipe

import (
	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/google/uuid"
)

// StepOption is a step creation option.
type StepOption func(*Step)

// WithStepID sets the step id.
func WithStepID(id uuid.UUID) StepOption {
	return func(s *Step) {
		s.id = id
	}
}

// Step is a recipe instruction.
type Step struct {
	id      uuid.UUID
	Details string
}

// ID returns the step id.
func (s *Step) ID() uuid.UUID {
	return s.id
}

// Validate the step.
func (s *Step) Validate() error {
	issues := []string{}

	if s.Details == "" {
		issues = append(issues, "details cannot be empty")
	}

	if len(issues) == 0 {
		return nil
	}

	return &entity.ValidationError{
		Issues: issues,
	}
}

// NewStep creates a new recipe step.
func NewStep(details string, opts ...StepOption) *Step {
	step := &Step{
		id:      uuid.New(),
		Details: details,
	}

	for _, opt := range opts {
		opt(step)
	}

	return step
}
