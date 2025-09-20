package store

import (
	"context"
	"database/sql"
	"social/internal/models"
)

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *models.User) error {
	query := `
	INSERT INTO users(
		username,
		email,
		password,
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
