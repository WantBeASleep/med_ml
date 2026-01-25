package card

import (
	"context"
	"errors"
	"net/http"

	"composition-api/internal/domain"
	med_domain "composition-api/internal/domain/med"
	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/mappers"
)

func (h *handler) MedCardPost(ctx context.Context, req *api.Card) (api.MedCardPostRes, error) {
	_, err := h.services.CardService.CreateCard(ctx, med_domain.Card{
		PatientID: req.PatientID,
		DoctorID:  req.DoctorID,
		Diagnosis: mappers.FromOptString(req.Diagnosis),
	})
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrBadRequest):
			return &api.MedCardPostBadRequest{
				StatusCode: http.StatusBadRequest,
				Response: api.Error{
					Message: "Неверный формат ОМС пациента",
				},
			}, nil
		case errors.Is(err, domain.ErrUnprocessableEntity):
			return &api.MedCardPostUnprocessableEntity{
				StatusCode: http.StatusUnprocessableEntity,
				Response: api.Error{
					Message: "Ошибка валидации данных",
				},
			}, nil
		default:
			return nil, err
		}
	}

	return &api.MedCardPostOK{}, nil
}
