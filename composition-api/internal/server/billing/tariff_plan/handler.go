package tariff_plan

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type TariffPlanHandler interface {
	TariffPlansIDGet(ctx context.Context, params api.TariffPlansIDGetParams) (api.TariffPlansIDGetRes, error)
	TariffPlansGet(ctx context.Context) (api.TariffPlansGetRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) TariffPlanHandler {
	return &handler{
		services: services,
	}
}
