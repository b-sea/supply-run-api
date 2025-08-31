package entity

import (
	"github.com/google/uuid"
	"github.com/lithammer/shortuuid"
	"github.com/zeebo/xxh3"
)

type ID struct {
	key string
}

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

func (id ID) IsValid() bool {
	return id.key != ""
}

func (id ID) String() string {
	return id.key
}
