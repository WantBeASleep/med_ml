package card

import (
	"context"

	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/med/mappers"

	"github.com/AlekSi/pointer"
)

func (h *handler) MedCardDoctorIDPatientIDGet(ctx context.Context, params api.MedCardDoctorIDPatientIDGetParams) (api.MedCardDoctorIDPatientIDGetRes, error) {
	card, err := h.services.CardService.GetCard(ctx, params.DoctorID, params.PatientID)
	if err != nil {
		return nil, err
	}

	return pointer.To(mappers.Card{}.Api(card)), nil
}
