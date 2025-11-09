package card

import (
	"context"
	"errors"

	"composition-api/internal/domain"
	med_domain "composition-api/internal/domain/med"
	api "composition-api/internal/generated/http/api"
	medmappers "composition-api/internal/server/med/mappers"

	"github.com/AlekSi/pointer"
)

func (h *handler) MedCardDoctorIDPatientIDPatch(ctx context.Context, req *api.MedCardDoctorIDPatientIDPatchReq, params api.MedCardDoctorIDPatientIDPatchParams) (api.MedCardDoctorIDPatientIDPatchRes, error) {
	card, err := h.services.CardService.UpdateCard(ctx, med_domain.Card{
		DoctorID:  params.DoctorID,
		PatientID: params.PatientID,
		Diagnosis: &req.Diagnosis,
	})
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return &api.MedCardDoctorIDPatientIDPatchNotFound{
				StatusCode: 404,
				Response: api.Error{
					Code:    404,
					Message: "Карта пациента не найдена",
				},
			}, nil
		case errors.Is(err, domain.ErrBadRequest):
			return &api.MedCardDoctorIDPatientIDPatchBadRequest{
				StatusCode: 400,
				Response: api.Error{
					Code:    400,
					Message: "Неверный формат запроса",
				},
			}, nil
		case errors.Is(err, domain.ErrUnprocessableEntity):
			return &api.MedCardDoctorIDPatientIDPatchUnprocessableEntity{
				StatusCode: 422,
				Response: api.Error{
					Code:    422,
					Message: "Ошибка валидации данных",
				},
			}, nil
		default:
			return nil, err
		}
	}
	return pointer.To(medmappers.Card{}.Api(card)), nil
}
