package repository

import (
	"context"
	"github.com/aidenismee/go-lambda-sample/internal/model"
)

type UserRepository struct {
	db any
}

func NewUserRepository(db any) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Get(ctx context.Context, id int) (*model.User, error) {
	return &model.User{
		ID:    1,
		Name:  "Mock User",
		Email: "mock.user@gmail.com",
	}, nil
}
