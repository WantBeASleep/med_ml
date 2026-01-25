package original_image

import (
	services "composition-api/internal/services"
)

type OriginalImageHandler interface {
	// Endpoints для original_image не определены в swagger.json
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) OriginalImageHandler {
	return &handler{
		services: services,
	}
}
