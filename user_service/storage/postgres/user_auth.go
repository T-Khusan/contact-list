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
func NewUserRepo(db *sqlx.DB) repo.UserRepoI {
	return &userRepo{db: db}
}

// CreateUser ...
func (r *userRepo) CreateUser(req *user_service.User) (string, error) {
	var id string
	query := "INSERT INTO users (name, lastname, password) VALUES ($1, $2, $3) RETURNING id"

	row := r.db.QueryRow(query, req.Name, req.Lastname, req.Password)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

// GetUser get database user
func (r *userRepo) GetUser(name, password string) (*user_service.User, error) {
	var user user_service.User
	query := "SELECT id FROM users WHERE name=$1 AND password=$2"
	err := r.db.Get(&user, query, name, password)

	return &user, err
}

/*
func (r *userRepo) GetUser(name, password string) (*user_service.GetTokenResponse, error) {
	var id user_service.GetTokenResponse
	query := "SELECT id FROM users WHERE name=$1 AND password=$2"
	err := r.db.Get(&id, query, name, password)

	return &id, err
}


*/
