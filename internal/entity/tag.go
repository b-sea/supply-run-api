// Package entity defines common data types.
package entity

import "strings"

// Tag is a simple string tag.
type Tag struct {
	name string
}

// NewTag creates a new Tag by name.
func NewTag(name string) Tag {
	return Tag{
		name: strings.ToLower(name),
	}
}

// ID returns the Tag ID.
func (t Tag) ID() ID {
	return NewSeededID([]byte(t.name))
}

// IsValid returns if the Tag is valid.
func (t Tag) IsValid() bool {
	return t.name != ""
}

func (t Tag) String() string {
	return t.name
}
