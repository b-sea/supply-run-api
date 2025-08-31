// Package entity defines common data types.
package entity

import (
	"github.com/google/uuid"
	"github.com/lithammer/shortuuid"
	"github.com/zeebo/xxh3"
)

// ID is a unique identifier.
type ID struct {
	key string
}

// NewID generates a new random ID.
func NewID() ID {
	return ID{
		key: shortuuid.New(),
	}
}

// NewSeededID generates a pre-determined ID from a seed.
func NewSeededID(seed []byte) ID {
	hash := xxh3.Hash128(seed).Bytes()
	key, _ := uuid.FromBytes(hash[:])

	return ID{
		key: shortuuid.DefaultEncoder.Encode(key),
	}
}

// IsValid returns if an ID is valid.
func (id ID) IsValid() bool {
	return id.key != ""
}

func (id ID) String() string {
	return id.key
}
