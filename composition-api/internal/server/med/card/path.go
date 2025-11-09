package card

import (
	"context"
	"errors"
	"net/http"

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
				StatusCode: http.StatusNotFound,
				Response: api.Error{
					Message: "Карта пациента не найдена",
				},
			}, nil
		case errors.Is(err, domain.ErrBadRequest):
			return &api.MedCardDoctorIDPatientIDPatchBadRequest{
				StatusCode: http.StatusBadRequest,
				Response: api.Error{
					Message: "Неверный формат запроса",
				},
			}, nil
		case errors.Is(err, domain.ErrUnprocessableEntity):
			return &api.MedCardDoctorIDPatientIDPatchUnprocessableEntity{
				StatusCode: http.StatusUnprocessableEntity,
				Response: api.Error{
					Message: "Ошибка валидации данных",
				},
			}, nil
		default:
			return nil, err
		}
	}
	return pointer.To(medmappers.Card{}.Api(card)), nil
}
