package repository

import (
	"context"
	"fmt"

	"example.com/fxdemo/api-gateway/domain"
	"github.com/jackc/pgx/v5"
)

type IRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByID(ctx context.Context, user *domain.User) (*domain.User, error)
}

type Repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) IRepository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	row := r.db.QueryRow(ctx, query, user.Name, user.Email)

	err := row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("could not create user: %w", err)
	}

	return user, nil
}

func (r *Repository) GetUserByID(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = $1`
	row := r.db.QueryRow(ctx, query, user.ID)

	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("could not get user: %w", err)
	}

	return user, nil
}
