package unit

import (
	"time"

	"github.com/google/uuid"
)

// Type is a base or derived SI type.
type Type struct {
	id        uuid.UUID
	createdAt time.Time
	updatedAt *time.Time
	owner     uuid.UUID
	name      string
}

// ID returns the id of the type.
func (t *Type) ID() uuid.UUID {
	return t.id
}

// CreatedAt returns a timestamp when the type is created.
func (t *Type) CreatedAt() time.Time {
	return t.createdAt
}

// UpdatedAt returns a timestamp when the type is updated.
func (t *Type) UpdatedAt() *time.Time {
	return t.updatedAt
}

// Owner returns the creator of the type.
func (t *Type) Owner() uuid.UUID {
	return t.owner
}

// Name returns the name of the type.
func (t *Type) Name() string {
	return t.name
}

func (t *Type) validate() error {
	issues := []string{}

	if t.name == "" {
		issues = append(issues, "name cannot be empty")
	}

	if len(issues) > 0 {
		return &ValidationError{
			Issues: issues,
		}
	}

	return nil
}

// UpdateTypeInput is used to update an existing type.
type UpdateTypeInput struct {
	Now  TimeFunc
	Name string
}

// Update and validate an existing type.
func (t *Type) Update(input UpdateTypeInput) error {
	if input.Now == nil {
		input.Now = time.Now
	}

	now := input.Now().UTC()
	t.updatedAt = &now
	t.name = input.Name

	if err := t.validate(); err != nil {
		return err
	}

	return nil
}

// NewTypeInput is used to create a new type.
type NewTypeInput struct {
	ID    UUIDFunc
	Now   TimeFunc
	Owner uuid.UUID
	Name  string
}

// NewType creates and validates a new type.
func NewType(input NewTypeInput) (*Type, error) {
	if input.ID == nil {
		input.ID = uuid.New
	}

	if input.Now == nil {
		input.Now = time.Now
	}

	result := &Type{
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
