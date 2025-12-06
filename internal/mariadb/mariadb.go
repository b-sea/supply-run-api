// Package mariadb implements all interactions with MariaDB.
package mariadb

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"path/filepath"
	"time"

	// Register the mysql driver.
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"

	"github.com/b-sea/go-server/server"
	"github.com/b-sea/supply-run-api/internal/query"
)

const defaultTimeout = 20 * time.Second

//go:embed sql/*.sql
var sqlFS embed.FS

type Recorder interface {
	ObserveMariaDBTxDuration(status string, duration time.Duration)
}

type Connector func() *sql.DB

func BasicConnector(host string, user string, pwd string) Connector {
	return func() *sql.DB {
		db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/?parseTime=true&multiStatements=true", user, pwd, host))

		return db
	}
}

var (
	_ server.HealthChecker   = (*Repository)(nil)
	_ query.RecipeRepository = (*Repository)(nil)
	_ query.UnitRepository   = (*Repository)(nil)
	_ query.UserRepository   = (*Repository)(nil)
)

type Repository struct {
	db       *sql.DB
	recorder Recorder
	timeout  time.Duration
}

func NewRepository(connector Connector, recorder Recorder) *Repository {
	return &Repository{
		db:       connector(),
		recorder: recorder,
		timeout:  defaultTimeout,
	}
}

func (r *Repository) HealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	if err := r.db.PingContext(ctx); err != nil {
		return databaseError(err)
	}

	return nil
}

func (r *Repository) Setup() error {
	var err error

	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	log := zerolog.Ctx(ctx)

	defer func() {
		event := log.Info()
		message := "setup complete"

		if err != nil {
			event = log.Error().Err(err)
			message = "setup failed"
		}

		duration := time.Since(start)

		event.Dur("duration_ms", duration).Msg(message)
	}()

	entries, err := sqlFS.ReadDir("sql")
	if err != nil {
		return fileReadError(err)
	}

	err = r.withTx(ctx, func(tx *sql.Tx) error {
		for _, entry := range entries {
			file := filepath.Join("sql", entry.Name())

			log.Debug().Str("file", file).Msg("loading file")

			cmd, err := sqlFS.ReadFile(file)
			if err != nil {
				return fileReadError(err)
			}

			if _, err := tx.ExecContext(ctx, string(cmd)); err != nil {
				return err //nolint: wrapcheck
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Close() error {
	if err := r.db.Close(); err != nil {
		return databaseError(err)
	}

	return nil
}

func (r *Repository) withTx(ctx context.Context, fn func(tx *sql.Tx) error) error {
	var err error

	start := time.Now()

	defer func() {
		log := zerolog.Ctx(ctx)

		event := log.Debug()
		status := "success"
		message := "transaction complete"

		if err != nil {
			event = log.Error().Err(err)
			status = "failed"
			message = "transaction failed"
		}

		duration := time.Since(start)

		event.Dur("duration_ms", duration).Msg(message)
		r.recorder.ObserveMariaDBTxDuration(status, duration)
	}()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return transactionError(err)
	}

	if err = fn(tx); err != nil {
		_ = tx.Rollback()

		return transactionError(err)
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()

		return transactionError(err)
	}

	return nil
}
