package postgres

import (
	"user_service"

	"github.com/jmoiron/sqlx"
)

// AuthPostgres ...
type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres ...
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

// CreateUser ...
func (r *AuthPostgres) CreateUser(user user_service.User) (int, error) {
	var id int
	query := "INSERT INTO users (name, username, password) VALUES ($1, $2, $3) RETURNING id"

	row := r.db.QueryRow(query, user.Name, user.Lastname, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
