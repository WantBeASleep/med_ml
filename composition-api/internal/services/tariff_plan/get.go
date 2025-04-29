package tariff_plan

import (
	"context"

	"github.com/google/uuid"

	domain "composition-api/internal/domain/billing"
)

func (s *service) GetTariffPlanByID(ctx context.Context, id uuid.UUID) (domain.TariffPlan, error) {
	return s.adapters.Billing.GetTariffPlanByID(ctx, id)
}

func (s *service) ListTariffPlans(ctx context.Context) ([]domain.TariffPlan, error) {
	return s.adapters.Billing.ListTariffPlans(ctx)
}
