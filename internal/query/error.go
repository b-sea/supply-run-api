package query

import (
	"errors"
	"fmt"
)

// ErrQuery is raised when a query fails.
var ErrQuery = errors.New("query error")

func queryError(err error) error {
	return fmt.Errorf("%w: %w", ErrQuery, err)
}
