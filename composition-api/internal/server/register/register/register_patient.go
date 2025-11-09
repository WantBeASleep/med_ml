package register

import (
	"context"
	"errors"
	"net/http"

	"composition-api/internal/domain"
	api "composition-api/internal/generated/http/api"
	"composition-api/internal/services/register"

	"github.com/AlekSi/pointer"
)

func (h *handler) RegPatientPost(ctx context.Context, req *api.RegPatientPostReq) (api.RegPatientPostRes, error) {
	id, err := h.services.RegisterService.RegisterPatient(ctx, register.RegisterPatientArg{
		Email:     req.Email,
		Password:  req.Password,
		FullName:  req.Fullname,
		Policy:    req.Policy,
		BirthDate: req.BirthDate,
	})
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrBadRequest):
			return &api.RegPatientPostBadRequest{
				StatusCode: http.StatusBadRequest,
				Response: api.Error{
					Message: "Неверный формат запроса",
				},
			}, nil
		case errors.Is(err, domain.ErrConflict):
			return &api.RegPatientPostConflict{
				StatusCode: http.StatusConflict,
				Response: api.Error{
					Message: "Пользователь с таким email уже существует",
				},
			}, nil
		case errors.Is(err, domain.ErrUnprocessableEntity):
			return &api.RegPatientPostUnprocessableEntity{
				StatusCode: http.StatusUnprocessableEntity,
				Response: api.Error{
					Message: "Ошибка валидации данных",
				},
			}, nil
		default:
			return nil, err
		}
	}

	return pointer.To(api.SimpleUuid{ID: id}), nil
}
