package service

import (
	"user_service"
)

// Authorization ...
type Authorization interface {
	CreateUser(user user_service.User) (int, error)
}

type Service struct {
	Authorization
}
