package store

import (
	"database/sql"
)

type Storage struct {
	Posts interface{ PostRepository }
	Users interface{ UserRepository }
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UsersStore{db},
	}
}
