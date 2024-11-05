package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"go.uber.org/fx"
)

// var Module = fx.Options(
// 	fx.Provide(NewPostgresConnection),
// )

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresConnection(lc fx.Lifecycle) (*pgx.Conn, error) {
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
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return conn.Ping(ctx)
		},
		OnStop: func(ctx context.Context) error {
			conn.Close(ctx)
			return nil
		},
	})

	return conn, nil
}
