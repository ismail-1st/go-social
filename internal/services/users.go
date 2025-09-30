package services

import (
	"context"
	"social/internal/models"
	"social/internal/store"
	"social/pkg/validation"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Store store.UserRepository
}

func (s UserService) Create(ctx context.Context, user *models.User) error {
	// validate domain model
	if err := validation.Validate.Struct(user); err != nil {
		return err
	}

	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)

	return s.Store.Create(ctx, user)
}
