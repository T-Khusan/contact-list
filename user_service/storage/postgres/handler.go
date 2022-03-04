package postgres

import "user_service/service"

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
