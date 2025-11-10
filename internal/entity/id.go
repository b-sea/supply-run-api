// Package entity defines all shared entities.
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

// NewID creates a new ID from a string.
func NewID(key string) ID {
	return ID{
		key: key,
	}
}

// NewRandomID generates a new random ID.
func NewRandomID() ID {
	return ID{
		key: shortuuid.New(),
	}
}

// NewSeededID generates a pre-determined ID from a seed.
func NewSeededID(seed string) ID {
	hash := xxh3.Hash128([]byte(seed)).Bytes()
	key, _ := uuid.FromBytes(hash[:])

	return ID{
		key: shortuuid.DefaultEncoder.Encode(key),
	}
}

func (id ID) String() string {
	return id.key
}
