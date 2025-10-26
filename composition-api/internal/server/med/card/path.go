package card

import (
	"context"
	"errors"

	adapter_errors "composition-api/internal/adapters/errors"
	domain "composition-api/internal/domain/med"
	api "composition-api/internal/generated/http/api"
	medmappers "composition-api/internal/server/med/mappers"

	"github.com/AlekSi/pointer"
)

func (h *handler) MedCardDoctorIDPatientIDPatch(ctx context.Context, req *api.MedCardDoctorIDPatientIDPatchReq, params api.MedCardDoctorIDPatientIDPatchParams) (api.MedCardDoctorIDPatientIDPatchRes, error) {
	card, err := h.services.CardService.UpdateCard(ctx, domain.Card{
		DoctorID:  params.DoctorID,
		PatientID: params.PatientID,
		Diagnosis: &req.Diagnosis,
	})
	if err != nil {
		switch {
		case errors.Is(err, adapter_errors.ErrNotFound):
			return &api.MedCardDoctorIDPatientIDPatchNotFound{
				StatusCode: 404,
				Response: api.Error{
					Code:    404,
					Message: "Карта не найдена",
				},
			}, nil
		default:
			return nil, err
		}
	}
	return pointer.To(medmappers.Card{}.Api(card)), nil
}
