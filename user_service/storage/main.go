package storage

import (
	"user_service"
	"user_service/service"
)

// Handler struct
type Handler struct {
	service *service.Service
}

// NewHandler ...
func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Authorization ...
type Authorization interface {
	CreateUser(user user_service.User) (int, error)
}

type Service struct {
	Authorization
}
