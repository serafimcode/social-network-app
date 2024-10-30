package services

import (
	"context"
	"log/slog"
	"social-network-app/internal/domain"
	"social-network-app/internal/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
	log            *slog.Logger
}

func BuildUserService(userRepository *repositories.UserRepository, log *slog.Logger) *UserService {
	return &UserService{userRepository: userRepository, log: log}
}

func (us *UserService) GetUsers(ctx context.Context) ([]domain.User, error) {
	users, err := us.userRepository.GetUsers(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}
