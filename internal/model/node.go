// Package model defines all data entities shared between front end, service, and repository layers.
package model

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type kind string

const (
	idSep = ":"
	idLen = 2
)

// ID is a compound global id made up of a storage key and a type.
type ID struct {
	Key  string
	Kind kind
}

// String encodes an ID into a string.
func (m ID) String() string {
	return base64.StdEncoding.EncodeToString([]byte(string(m.Kind) + idSep + m.Key))
}

// FromString decodes an ID string into an ID.
func (m *ID) FromString(encoded string) error {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	components := strings.Split(string(decoded), idSep)
	if len(components) != idLen {
		return fmt.Errorf("incorrect id format") //nolint: goerr113
	}

	m.Key = components[1]
	m.Kind = kind(components[0])

	return nil
}

// MarshalJSON converts an ID into a JSON representation.
func (m *ID) MarshalJSON() ([]byte, error) {
	result, err := json.Marshal(m.String())
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return result, nil
}

// UnmarshalJSON converts a JSON ID representation into an ID.
func (m ID) UnmarshalJSON(data []byte) error {
	if err := m.FromString(string(data)); err != nil {
		return err
	}
	return nil
}

// Node defines all functions required on an Entity.
type Node interface {
	GetID() ID
}

type CreateInput[N Node] interface {
	Validate() error
	ToEntity(key string, timestamp time.Time) N
}

type UpdateInput[N Node] interface {
	Validate() error
	MergeEntity(N *Node, timestamp time.Time) N
}
