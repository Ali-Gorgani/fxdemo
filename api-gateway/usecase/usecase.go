package usecase

import (
	"context"
	"errors"

	"example.com/fxdemo/api-gateway/domain"
	"example.com/fxdemo/api-gateway/repository"
)

type IUsecase interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByID(ctx context.Context, user *domain.User) (*domain.User, error)
}

type Usecase struct {
	repo repository.IRepository
}

func NewUsecase(repo repository.IRepository) IUsecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	if user.Email == "" {
		return nil, errors.New("email is required")
	}
	return u.repo.CreateUser(ctx, user)
}

func (u *Usecase) GetUserByID(ctx context.Context, user *domain.User) (*domain.User, error) {
	foundUser, err := u.repo.GetUserByID(ctx, user)
	if err != nil {
		return nil, err
	}
	return foundUser, nil
}
