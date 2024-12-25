package services

import (
	"context"

	"github.com/KalinduGandara/erp-system/internal/domain/entities"
	"github.com/KalinduGandara/erp-system/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, username, password, role string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entities.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}

	return s.userRepo.Create(ctx, user)
}

func (s *UserService) ValidateUser(ctx context.Context, username, password string) bool {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
