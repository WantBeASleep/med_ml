package card

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type CardHandler interface {
	MedCardPost(ctx context.Context, req *api.Card) (api.MedCardPostRes, error)
	MedCardDoctorIDPatientIDGet(ctx context.Context, params api.MedCardDoctorIDPatientIDGetParams) (api.MedCardDoctorIDPatientIDGetRes, error)
	MedCardDoctorIDPatientIDPatch(ctx context.Context, req *api.MedCardDoctorIDPatientIDPatchReq, params api.MedCardDoctorIDPatientIDPatchParams) (api.MedCardDoctorIDPatientIDPatchRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) CardHandler {
	return &handler{
		services: services,
	}
}
