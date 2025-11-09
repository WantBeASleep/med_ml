package card

import (
	"context"
	"errors"
	"net/http"

	"composition-api/internal/domain"
	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/med/mappers"

	"github.com/AlekSi/pointer"
)

func (h *handler) MedCardDoctorIDPatientIDGet(ctx context.Context, params api.MedCardDoctorIDPatientIDGetParams) (api.MedCardDoctorIDPatientIDGetRes, error) {
	card, err := h.services.CardService.GetCard(ctx, params.DoctorID, params.PatientID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return &api.MedCardDoctorIDPatientIDGetNotFound{
				StatusCode: http.StatusNotFound,
				Response: api.Error{
					Message: "Карта пациента не найдена",
				},
			}, nil
		default:
			return nil, err
		}
	}

	return pointer.To(mappers.Card{}.Api(card)), nil
}
