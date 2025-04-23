package service

import (
	"context"
	"github.com/aidenismee/go-lambda-sample/internal/model"
)

type UserService struct {
	userRepository IUserRepository
}

type IUserRepository interface {
	Get(ctx context.Context, id int) (*model.User, error)
}

func NewUserService(userRepository IUserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) GetDetailUser(ctx context.Context, id int) (*model.User, error) {
	return s.userRepository.Get(ctx, id)
}
