package entity

import "strings"

type Tag struct {
	name string
}

func NewTag(name string) Tag {
	return Tag{
		name: strings.ToLower(name),
	}
}

func (t Tag) String() string {
	return t.name
}

func (t Tag) ID() ID {
	return NewSeededID([]byte(t.name))
}

func (t Tag) IsValid() bool {
	return t.name != ""
}
