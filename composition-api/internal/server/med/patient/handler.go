package patient

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type PatientHandler interface {
	MedPatientPost(ctx context.Context, req *api.MedPatientPostReq) (api.MedPatientPostRes, error)
	MedPatientIDGet(ctx context.Context, params api.MedPatientIDGetParams) (api.MedPatientIDGetRes, error)
	MedDoctorIDPatientsGet(ctx context.Context, params api.MedDoctorIDPatientsGetParams) (api.MedDoctorIDPatientsGetRes, error)
	MedPatientIDPatch(ctx context.Context, req *api.MedPatientIDPatchReq, params api.MedPatientIDPatchParams) (api.MedPatientIDPatchRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) PatientHandler {
	return &handler{
		services: services,
	}
}
