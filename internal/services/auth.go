package services

import (
	"context"
	"fmt"
	"social/internal/dto"
	"social/internal/models"
	"social/internal/store"
	"social/pkg/auth"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Store store.UserRepository
}

func (s AuthService) Register(ctx context.Context, user *dto.UserRegister) (*models.User, error) {

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

func (s AuthService) Login(ctx context.Context, user *dto.UserLogin) (*dto.UserLoginResponse, error) {
	userData, err := s.Store.Login(ctx, &dto.UserLoginDetail{
		Email: user.Email,
	})

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	accessToken, err := auth.GenerateJwtToken(userData.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := auth.GenerateRefreshToken(userData.ID)
	if err != nil {
		return nil, err
	}

	return &dto.UserLoginResponse{
		ID:           userData.ID,
		Email:        userData.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
