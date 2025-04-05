package doctor

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type DoctorHandler interface {
	MedDoctorIDGet(ctx context.Context, params api.MedDoctorIDGetParams) (api.MedDoctorIDGetRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) DoctorHandler {
	return &handler{
		services: services,
	}
}
