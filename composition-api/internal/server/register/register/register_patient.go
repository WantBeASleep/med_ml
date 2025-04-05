package register

import (
	"context"

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
		return nil, err
	}

	return pointer.To(api.SimpleUuid{ID: id}), nil
}
