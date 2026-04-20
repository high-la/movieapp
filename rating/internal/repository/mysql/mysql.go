package mysql

import "database/sql"

// Repository defines a MySQL-based rating repository.
type Repository struct {
	db *sql.DB
}

// New creates a new MySQL-based rating repository.
func New() (*Repository, error) {
	db, err := sql.Open("mysql", "root:1234@/movieexample")
	if err != nil {
		return nil, err
	}
	return &Repository{db}, nil
}
