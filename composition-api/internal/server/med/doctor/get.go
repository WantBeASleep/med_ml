package doctor

import (
	"context"

	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/med/mappers"

	"github.com/AlekSi/pointer"
)

func (h *handler) MedDoctorIDGet(ctx context.Context, params api.MedDoctorIDGetParams) (api.MedDoctorIDGetRes, error) {
	doctor, err := h.services.DoctorService.GetDoctor(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	return pointer.To(mappers.Doctor{}.Api(doctor)), nil
}
