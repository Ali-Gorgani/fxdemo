package repository

import (
	"context"
	"fmt"

	"example.com/fxdemo/api-gateway/domain"
	"example.com/fxdemo/ent"
	"example.com/fxdemo/ent/user"
)

type EntRepository struct {
	client *ent.Client
}

func NewEntRepository(client *ent.Client) IRepository {
	return &EntRepository{client: client}
}

func (r *EntRepository) CreateUser(ctx context.Context, u *domain.User) (*domain.User, error) {
	createdUser, err := r.client.User.Create().SetName(u.Name).SetEmail(u.Email).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not create user: %w", err)
	}

	u.ID = createdUser.ID
	return u, nil
}

func (r *EntRepository) GetUserByID(ctx context.Context, u *domain.User) (*domain.User, error) {
	foundUser, err := r.client.User.Query().Where(user.IDEQ(u.ID)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get user: %w", err)
	}

	u.Name = foundUser.Name
	u.Email = foundUser.Email
	return u, nil
}
