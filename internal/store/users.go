package store

import (
	"context"
	"database/sql"
	"fmt"
	"social/internal/dto"
	"social/internal/models"
)

type UserRepository interface {
	Create(context.Context, *models.User) error
	Login(context.Context, *dto.UserLoginDetail) (*dto.UserLoginDetail, error)
}

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *models.User) error {
	query := `
	INSERT INTO users(
		username,
		email,
		password
	) VALUES (
	 	$1,
		$2,
		$3
	 ) RETURNING id, created_at, updated_at
	 `

	err := s.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *UsersStore) Login(ctx context.Context, user *dto.UserLoginDetail) (*dto.UserLoginDetail, error) {
	query := `
	SELECT id, email, password
	FROM users
	WHERE email = $1
	`

	var resp dto.UserLoginDetail

	err := s.db.QueryRowContext(ctx, query, user.Email).Scan(&resp.ID, &resp.Email, &resp.Password)

	if err == sql.ErrNoRows {
		// Instead of returning DB error, signal invalid credentials
		return nil, fmt.Errorf("invalid credentials")
	}

	if err != nil {
		return nil, err
	}

	return &resp, nil
}
