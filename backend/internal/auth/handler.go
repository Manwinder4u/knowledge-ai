package auth

type Handler struct {
	service *Service
}

// Hanlder of Auth -- Which return the Service
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}
