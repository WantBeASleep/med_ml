package tariff_plan

import (
	"context"

	api "composition-api/internal/generated/http/api"

	"github.com/AlekSi/pointer"
)

func (h *handler) TariffPlansIDGet(ctx context.Context, params api.TariffPlansIDGetParams) (api.TariffPlansIDGetRes, error) {
	tariffPlan, err := h.services.TariffPlanService.GetTariffPlanByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	return pointer.To(api.TariffPlan{
		ID:          tariffPlan.ID,
		Name:        tariffPlan.Name,
		Description: tariffPlan.Description,
		Price:       tariffPlan.Price.String(),
		Duration:    int(tariffPlan.Duration.Seconds()),
	}), nil
}

func (h *handler) TariffPlansGet(ctx context.Context) (api.TariffPlansGetRes, error) {
	tariffPlans, err := h.services.TariffPlanService.ListTariffPlans(ctx)
	if err != nil {
		return nil, err
	}

	var result []api.TariffPlan
	for _, tp := range tariffPlans {
		result = append(result, api.TariffPlan{
			ID:          tp.ID,
			Name:        tp.Name,
			Description: tp.Description,
			Price:       tp.Price.String(),
			Duration:    int(tp.Duration.Seconds()),
		})
	}

	return pointer.To(api.TariffPlansGetOKApplicationJSON(result)), nil
}
