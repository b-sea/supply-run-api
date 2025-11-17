package model

import (
	"encoding/base64"
	"io"
	"strconv"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/b-sea/supply-run-api/internal/entity"
)

// Kind is an internal data type that is used in IDs.
type Kind string

// RecipeKind, et al. are custom GraphQL relay types.
const (
	RecipeKind = Kind("recipe")
	UserKind   = Kind("user")

	delim            = ":"
	idSplitCount     = 2
	cursorSplitCount = 2
)

// ID is a global identifier.
type ID struct {
	Key  entity.ID
	Kind Kind
}

func (s ID) String() string {
	return base64.StdEncoding.EncodeToString([]byte(string(s.Kind) + delim + s.Key.String()))
}

// MarshalID marshals an ID to a GraphQL format.
func MarshalID(value ID) graphql.Marshaler { //nolint: ireturn
	return graphql.WriterFunc(
		func(writer io.Writer) {
			_, _ = io.WriteString(writer, strconv.Quote(value.String()))
		},
	)
}

// UnmarshalID unmarshals a value into an ID.
func UnmarshalID(value any) (ID, error) {
	str, ok := value.(string)
	if !ok {
		return ID{}, nil
	}

	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ID{}, nil //nolint: nilerr
	}

	split := strings.Split(string(decoded), delim)
	if len(split) != idSplitCount {
		return ID{}, nil
	}

	return ID{
		Key:  entity.NewID(split[1]),
		Kind: Kind(split[0]),
	}, nil
}

// Cursor is placement of an item on a page.
type Cursor struct {
	ID   entity.ID
	Sort Sort
}

func (s Cursor) String() string {
	return base64.StdEncoding.EncodeToString([]byte(s.ID.String() + delim + s.Sort.String()))
}

// MarshalCursor marshals a Cursor to a GraphQL format.
func MarshalCursor(value Cursor) graphql.Marshaler { //nolint: ireturn
	return graphql.WriterFunc(
		func(writer io.Writer) {
			_, _ = io.WriteString(writer, strconv.Quote(value.String()))
		},
	)
}

// UnmarshalCursor unmarshals a value into a Cursor.
func UnmarshalCursor(value any) (Cursor, error) {
	str, ok := value.(string)
	if !ok {
		return Cursor{}, nil
	}

	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return Cursor{}, nil //nolint: nilerr
	}

	split := strings.Split(string(decoded), delim)
	if len(split) != cursorSplitCount {
		return Cursor{}, nil
	}

	result := Cursor{
		ID:   entity.NewID(split[0]),
		Sort: Sort(split[1]),
	}

	return result, nil
}
