package card

import (
	"context"

	domain "composition-api/internal/domain/med"
	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/mappers"
)

func (h *handler) MedCardPost(ctx context.Context, req *api.Card) (api.MedCardPostRes, error) {
	err := h.services.CardService.CreateCard(ctx, domain.Card{
		PatientID: req.PatientID,
		DoctorID:  req.DoctorID,
		Diagnosis: mappers.FromOptString(req.Diagnosis),
	})
	if err != nil {
		return nil, err
	}

	return &api.MedCardPostOK{}, nil
}
