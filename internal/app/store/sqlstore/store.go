package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Store ...
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		s,
	}

	return s.userRepository
}
