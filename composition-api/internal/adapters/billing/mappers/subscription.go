package mappers

import (
	"github.com/google/uuid"

	domain "composition-api/internal/domain/billing"
	pb "composition-api/internal/generated/grpc/clients/billing"
)

type Subscription struct{}

func (m Subscription) Domain(pb *pb.Subscription) domain.Subscription {
	return domain.Subscription{
		ID:           uuid.MustParse(pb.SubscriptionId),
		TariffPlanID: uuid.MustParse(pb.TariffPlanId),
		Status:       domain.SubscriptionStatus(pb.Status.String()),
		StartDate:    pb.StartDate.AsTime(),
		EndDate:      pb.EndDate.AsTime(),
	}
}
