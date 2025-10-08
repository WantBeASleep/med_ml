package patient

import (
	"context"
	"errors"

	adapter_errors "composition-api/internal/adapters/errors"
	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/med/mappers"

	"github.com/AlekSi/pointer"
)

func (h *handler) MedPatientIDGet(ctx context.Context, params api.MedPatientIDGetParams) (api.MedPatientIDGetRes, error) {
	patient, err := h.services.PatientService.GetPatient(ctx, params.ID)
	if err != nil {
		switch {
		case errors.Is(err, adapter_errors.ErrNotFound):
			return &api.MedPatientIDGetNotFound{
				StatusCode: 404,
				Response: api.Error{
					Code:    404,
					Message: err.Error(),
				},
			}, nil
		default:
			return nil, err
		}
	}

	return pointer.To(mappers.Patient{}.Api(patient)), nil
}

func (h *handler) MedDoctorIDPatientsGet(ctx context.Context, params api.MedDoctorIDPatientsGetParams) (api.MedDoctorIDPatientsGetRes, error) {
	patients, err := h.services.PatientService.GetPatientsByDoctorID(ctx, params.ID)
	if err != nil && !errors.Is(err, adapter_errors.ErrNotFound) {
		return nil, err
	}

	return pointer.To(
		api.MedDoctorIDPatientsGetOKApplicationJSON(
			mappers.Patient{}.SliceApi(patients),
		),
	), nil
}
