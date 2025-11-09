package patient

import (
	"context"
	"errors"
	"net/http"

	"composition-api/internal/domain"
	api "composition-api/internal/generated/http/api"
	apimappers "composition-api/internal/server/mappers"
	"composition-api/internal/server/med/mappers"

	"composition-api/internal/services/patient"

	"github.com/AlekSi/pointer"
)

func (h *handler) MedPatientIDPatch(ctx context.Context, req *api.MedPatientIDPatchReq, params api.MedPatientIDPatchParams) (api.MedPatientIDPatchRes, error) {
	patient, err := h.services.PatientService.UpdatePatient(ctx, params.ID, patient.UpdatePatientArg{
		Active:     apimappers.FromOptBool(req.Active),
		Malignancy: apimappers.FromOptBool(req.Malignancy),
	})
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return &api.MedPatientIDPatchNotFound{
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
