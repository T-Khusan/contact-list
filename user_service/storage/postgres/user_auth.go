package postgres

import (
	"user_service/genproto/user_service"
	"user_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

// AuthPostgres ...
type userRepo struct {
	db *sqlx.DB
}

// NewAuthPostgres ...
func NewContactRepo(db *sqlx.DB) repo.UserRepoI {
	return &userRepo{db: db}
}

// CreateUser ...
func (r *userRepo) CreateUser(req *user_service.User) (string, error) {
	var id string
	query := "INSERT INTO users (name, username, password) VALUES ($1, $2, $3) RETURNING id"

	row := r.db.QueryRow(query, req.Name, req.Lastname, req.Password)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}
