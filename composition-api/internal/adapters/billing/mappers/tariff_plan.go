package mappers

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	domain "composition-api/internal/domain/billing"
	pb "composition-api/internal/generated/grpc/clients/billing"
)

type TariffPlan struct{}

func (m TariffPlan) Domain(pb *pb.TariffPlan) domain.TariffPlan {
	price, _ := decimal.NewFromString(pb.Price)
	return domain.TariffPlan{
		ID:          uuid.MustParse(pb.TariffPlanId),
		Name:        pb.Name,
		Description: pb.Description,
		Price:       price,
		Duration:    pb.Duration.AsDuration(),
	}
}

func (m TariffPlan) SliceDomain(pbs []*pb.TariffPlan) []domain.TariffPlan {
	var result []domain.TariffPlan
	for _, pb := range pbs {
		result = append(result, m.Domain(pb))
	}
	return result
}
