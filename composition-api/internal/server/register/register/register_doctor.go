package register

import (
	"context"

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
		return nil, err
	}

	return pointer.To(api.SimpleUuid{ID: id}), nil
}
