package services

import (
	"context"
	"social/internal/dto"
	"social/internal/models"
	"social/internal/store"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Store store.UserRepository
}

func (s UserService) Create(ctx context.Context, user *dto.UserRegister) (*models.User, error) {

	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: string(hashed),
	}

	if err := s.Store.Create(ctx, &u); err != nil {
		return nil, err
	}

	return &u, nil
}
