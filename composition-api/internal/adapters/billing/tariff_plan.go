package billing

import (
	"context"
	"fmt"

	"composition-api/internal/adapters/billing/mappers"
	domain "composition-api/internal/domain/billing"
	pb "composition-api/internal/generated/grpc/clients/billing"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/google/uuid"
)

func (a *adapter) GetTariffPlanByID(ctx context.Context, id uuid.UUID) (domain.TariffPlan, error) {
	res, err := a.client.GetTariffPlanByID(ctx, &pb.GetTariffPlanByIDIn{Id: id.String()})
	if err != nil {
		return domain.TariffPlan{}, fmt.Errorf("failed to get tariff plan by ID: %w", err)
	}
	return mappers.TariffPlan{}.Domain(res.TariffPlan), nil
}

func (a *adapter) ListTariffPlans(ctx context.Context) ([]domain.TariffPlan, error) {
	res, err := a.client.ListTariffPlans(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to list tariff plans: %w", err)
	}
	return mappers.TariffPlan{}.SliceDomain(res.TariffPlans), nil
}
