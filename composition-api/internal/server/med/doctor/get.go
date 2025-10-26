package doctor

import (
	"context"
	"errors"

	adapter_errors "composition-api/internal/adapters/errors"
	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/med/mappers"

	"github.com/AlekSi/pointer"
)

func (h *handler) MedDoctorIDGet(ctx context.Context, params api.MedDoctorIDGetParams) (api.MedDoctorIDGetRes, error) {
	doctor, err := h.services.DoctorService.GetDoctor(ctx, params.ID)
	if err != nil {
		switch {
		case errors.Is(err, adapter_errors.ErrNotFound):
			return &api.MedDoctorIDGetNotFound{
				StatusCode: 404,
				Response: api.Error{
					Code:    404,
					Message: "Врач не найден",
				},
			}, nil
		default:
			return nil, err
		}
	}

	return pointer.To(mappers.Doctor{}.Api(doctor)), nil
}
