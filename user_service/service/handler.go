package service

// Handler struct
type Handler struct {
	service *Service
}

// NewHandler ...
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}
