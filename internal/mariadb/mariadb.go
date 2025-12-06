package mariadb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	// Register the mysql driver.
	_ "github.com/go-sql-driver/mysql"

	"github.com/b-sea/go-server/server"
	"github.com/b-sea/supply-run-api/internal/entity"
	"github.com/b-sea/supply-run-api/internal/query"
)

const defaultTimeout = 20 * time.Second

type Recorder interface{}

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
	db      *sql.DB
	timeout time.Duration
}

func NewRepository(connector Connector) *Repository {
	return &Repository{
		db:      connector(),
		timeout: defaultTimeout,
	}
}

func (r *Repository) HealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	if err := r.db.PingContext(ctx); err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindRecipes(
	ctx context.Context,
	filter query.RecipeFilter,
	page query.Pagination,
	order query.Order,
) ([]*query.Recipe, error) {
	return nil, nil
}

func (r *Repository) GetRecipes(ctx context.Context, ids []entity.ID) ([]*query.Recipe, error) {
	return nil, nil
}

func (r *Repository) FindTags(ctx context.Context, filter *string) ([]string, error) {
	return nil, nil
}

func (r *Repository) GetUnits(ctx context.Context, ids []entity.ID) ([]*query.Unit, error) {
	return nil, nil
}

func (r *Repository) GetUsers(ctx context.Context, ids []entity.ID) ([]*query.User, error) {
	return nil, nil
}
