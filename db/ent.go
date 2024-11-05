package db

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"example.com/fxdemo/ent"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewEntClient),
)

// NewEntClient creates a new ent client.
func NewEntClient(lc fx.Lifecycle) (*ent.Client, error) {
	cfg := Config{
		Host:     "localhost",
		Port:     5432,
		User:     "root",
		Password: "secret",
		DBName:   "psql_db",
		SSLMode:  "disable",
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	driver, err := sql.Open(dialect.Postgres, dsn)
	if err != nil {
		return nil, err
	}
	client := ent.NewClient(ent.Driver(driver))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			client.Close()
			return nil
		},
	})

	return client, nil
}
