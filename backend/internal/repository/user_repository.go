package repository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/shuakr/AnimeConnect/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id string) (*domain.User, error)
}

type userRepository struct {
	client *firestore.Client
}

func NewUserRepository(client *firestore.Client) UserRepository {
	return &userRepository{client: client}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	user.CreatedAt = time.Now()
	_, err := r.client.Collection("users").Doc(user.ID).Set(ctx, user)

	return err
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	doc, err := r.client.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var user domain.User
	if err := doc.DataTo(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
