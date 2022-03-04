package repo

import (
	"user_service"
)

// Authorization ...
type Authorization interface {
	CreateUser(user user_service.User) (int, error)
	GetUser(username, password string) (user_service.User, error)
}

// Repository ...
type Repository struct {
	Authorization
}
