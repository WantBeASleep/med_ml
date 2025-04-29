package tariff_plan

import (
	"context"

	"github.com/google/uuid"

	"composition-api/internal/adapters"
	domain "composition-api/internal/domain/billing"
)

type Service interface {
	GetTariffPlanByID(ctx context.Context, id uuid.UUID) (domain.TariffPlan, error)
	ListTariffPlans(ctx context.Context) ([]domain.TariffPlan, error)
}

type service struct {
	adapters *adapters.Adapters
}

func New(adapters *adapters.Adapters) Service {
	return &service{
		adapters: adapters,
	}
}
