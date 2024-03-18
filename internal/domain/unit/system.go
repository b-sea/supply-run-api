package unit

import (
	"time"

	"github.com/google/uuid"
)

// System is a measurement system.
type System struct {
	id        uuid.UUID
	createdAt time.Time
	updatedAt *time.Time
	owner     uuid.UUID
	name      string
}

// ID returns the id of the system.
func (s *System) ID() uuid.UUID {
	return s.id
}

// CreatedAt returns a timestamp when the system was created.
func (s *System) CreatedAt() time.Time {
	return s.createdAt
}

// UpdatedAt returns a timestamp when the system was updated.
func (s *System) UpdatedAt() *time.Time {
	return s.updatedAt
}

// Owner returns the creator of the system.
func (s *System) Owner() uuid.UUID {
	return s.owner
}

// Name returns the name of the system.
func (s *System) Name() string {
	return s.name
}

func (s *System) validate() error {
	issues := []string{}

	if s.name == "" {
		issues = append(issues, "name cannot be empty")
	}

	if len(issues) > 0 {
		return &ValidationError{
			Issues: issues,
		}
	}

	return nil
}

// UpdateSystemInput is used to update an existing system.
type UpdateSystemInput struct {
	Now  TimeFunc
	Name string
}

// Update and validate an existing system.
func (s *System) Update(input UpdateSystemInput) error {
	if input.Now == nil {
		input.Now = time.Now
	}

	now := input.Now().UTC()
	s.updatedAt = &now
	s.name = input.Name

	if err := s.validate(); err != nil {
		return err
	}

	return nil
}

// NewSystemInput is used to create a new system.
type NewSystemInput struct {
	ID    UUIDFunc
	Now   TimeFunc
	Owner uuid.UUID
	Name  string
}

// NewSystem creates and validates a new system.
func NewSystem(input NewSystemInput) (*System, error) {
	if input.ID == nil {
		input.ID = uuid.New
	}

	if input.Now == nil {
		input.Now = time.Now
	}

	result := &System{
		id:        input.ID(),
		createdAt: input.Now().UTC(),
		owner:     input.Owner,
		name:      input.Name,
	}

	if err := result.validate(); err != nil {
		return nil, err
	}

	return result, nil
}
