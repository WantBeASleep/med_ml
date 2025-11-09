package patient

import (
	"context"
	"errors"
	"net/http"

	"composition-api/internal/domain"
	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/med/mappers"

	"github.com/AlekSi/pointer"
)

func (h *handler) MedPatientIDGet(ctx context.Context, params api.MedPatientIDGetParams) (api.MedPatientIDGetRes, error) {
	patient, err := h.services.PatientService.GetPatient(ctx, params.ID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return &api.MedPatientIDGetNotFound{
				StatusCode: http.StatusNotFound,
				Response: api.Error{
					Message: "Пациент не найден",
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
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return &api.MedDoctorIDPatientsGetNotFound{
				StatusCode: http.StatusNotFound,
				Response: api.Error{
					Message: "Врач не найден",
				},
			}, nil
		}
		return nil, err
	}

	return pointer.To(
		api.MedDoctorIDPatientsGetOKApplicationJSON(
			mappers.Patient{}.SliceApi(patients),
		),
	), nil
}
