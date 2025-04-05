package register

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type RegisterHandler interface {
	RegDoctorPost(ctx context.Context, req *api.RegDoctorPostReq) (api.RegDoctorPostRes, error)
	RegPatientPost(ctx context.Context, req *api.RegPatientPostReq) (api.RegPatientPostRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) RegisterHandler {
	return &handler{
		services: services,
	}
}
