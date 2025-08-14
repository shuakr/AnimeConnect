package service

import (
	"context"
	"errors"

	"github.com/shuakr/AnimeConnect/internal/domain"
	"github.com/shuakr/AnimeConnect/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, username, email, password string) (*domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(ctx context.Context, username, email, password string) (*domain.User, error) {
	if username == "" || email == "" || password == "" {
		return nil, errors.New("All fields are required")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		ID:       email,
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}
	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}
