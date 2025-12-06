package mariadb

import (
	"errors"
	"fmt"
)

// ErrDatabase, et al. are custom errors for MariaDB.
var (
	ErrDatabase    = errors.New("database error")
	ErrTransaction = errors.New("transaction error")
	ErrFileRead    = errors.New("file read error")
)

func databaseError(err error) error {
	return fmt.Errorf("%w: %w", ErrDatabase, err)
}

func transactionError(err error) error {
	return fmt.Errorf("%w: %w", ErrTransaction, err)
}

func fileReadError(err error) error {
	return fmt.Errorf("%w: %w", ErrFileRead, err)
}
