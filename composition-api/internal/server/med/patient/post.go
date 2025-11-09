package patient

import (
	"context"
	"errors"
	"net/http"

	"composition-api/internal/domain"
	api "composition-api/internal/generated/http/api"
	"composition-api/internal/services/patient"

	"github.com/AlekSi/pointer"
)

func (h *handler) MedPatientPost(ctx context.Context, req *api.MedPatientPostReq) (api.MedPatientPostRes, error) {
	id, err := h.services.PatientService.CreatePatient(ctx, patient.CreatePatientArg{
		Fullname:   req.Fullname,
		Email:      req.Email,
		Policy:     req.Policy,
		Active:     req.Active,
		Malignancy: req.Malignancy,
		BirthDate:  req.BirthDate,
	})
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrBadRequest):
			return &api.MedPatientPostBadRequest{
				StatusCode: http.StatusBadRequest,
				Response: api.Error{
					Message: "Неверный формат ОМС",
				},
			}, nil
		case errors.Is(err, domain.ErrUnprocessableEntity):
			return &api.MedPatientPostUnprocessableEntity{
				StatusCode: http.StatusUnprocessableEntity,
				Response: api.Error{
					Message: "Ошибка валидации данных",
				},
			}, nil
		case errors.Is(err, domain.ErrConflict):
			return &api.MedPatientPostConflict{
				StatusCode: http.StatusConflict,
				Response: api.Error{
					Message: "Пользователь с таким email уже существует",
				},
			}, nil
		default:
			return nil, err
		}
	}

	return pointer.To(api.SimpleUuid{ID: id}), nil
}
