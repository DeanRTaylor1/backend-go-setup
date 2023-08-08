package db

import (
	"database/sql"
)

type Store interface {
	Querier
}

// Store is a wrapper around the database connection that provides
// a transactional API for executing queries.
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
