// Package recipe defines everything to manage the recipes domain.
package recipe

import "github.com/google/uuid"

// TagOption is a tag creation option.
type TagOption func(*Tag)

// WithTagID sets the tag id.
func WithTagID(id uuid.UUID) TagOption {
	return func(t *Tag) {
		t.id = id
	}
}

// Tag is an organizing tag for recipes.
type Tag struct {
	id   uuid.UUID
	name string
}

// ID returns the tag id.
func (t *Tag) ID() uuid.UUID {
	return t.id
}

// Name returns the tag name.
func (t *Tag) Name() string {
	return t.name
}

// Validate the tag.
func (t *Tag) Validate() error {
	issues := []string{}

	if t.name == "" {
		issues = append(issues, "name cannot be empty")
	}

	if len(issues) == 0 {
		return nil
	}

	return &ValidationError{
		Issues: issues,
	}
}

// NewTag creates a new recipe tag.
func NewTag(name string, opts ...TagOption) *Tag {
	tag := &Tag{
		id:   uuid.New(),
		name: name,
	}

	for _, opt := range opts {
		opt(tag)
	}

	return tag
}
