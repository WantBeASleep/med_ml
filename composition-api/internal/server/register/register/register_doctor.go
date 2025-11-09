package register

import (
	"context"
	"errors"

	"composition-api/internal/domain"
	api "composition-api/internal/generated/http/api"
	apimappers "composition-api/internal/server/mappers"
	"composition-api/internal/services/register"

	"github.com/AlekSi/pointer"
)

func (h *handler) RegDoctorPost(ctx context.Context, req *api.RegDoctorPostReq) (api.RegDoctorPostRes, error) {
	id, err := h.services.RegisterService.RegisterDoctor(ctx, register.RegisterDoctorArg{
		Email:       req.Email,
		Password:    req.Password,
		FullName:    req.Fullname,
		Org:         req.Org,
		Job:         req.Job,
		Description: apimappers.FromOptString(req.Description),
	})
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrBadRequest):
			return &api.RegDoctorPostBadRequest{
				StatusCode: 400,
				Response: api.Error{
					Message: "Неверный формат запроса",
				},
			}, nil
		case errors.Is(err, domain.ErrConflict):
			return &api.RegDoctorPostConflict{
				StatusCode: 409,
				Response: api.Error{
					Message: "Пользователь с таким email уже существует",
				},
			}, nil
		case errors.Is(err, domain.ErrUnprocessableEntity):
			return &api.RegDoctorPostUnprocessableEntity{
				StatusCode: 422,
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
