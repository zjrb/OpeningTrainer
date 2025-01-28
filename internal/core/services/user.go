package services

import (
	"github.com/zjrb/OpeningTrainer/internal/core/domain"
	"github.com/zjrb/OpeningTrainer/internal/core/ports"
)

type UserService struct {
	userRepo ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUserByEmail(email string) (*domain.User, error) {
	return s.userRepo.GetUserByEmail(email)
}
