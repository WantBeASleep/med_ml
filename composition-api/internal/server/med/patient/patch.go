package patient

import (
	"context"

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
		return nil, err
	}

	return pointer.To(mappers.Patient{}.Api(patient)), nil
}
