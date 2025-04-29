package mappers

import (
	"github.com/google/uuid"

	domain "composition-api/internal/domain/billing"
	pb "composition-api/internal/generated/grpc/clients/billing"
)

type PurchaseSubscription struct{}

func (m PurchaseSubscription) Domain(pb *pb.PurchaseSubscriptionOut) domain.PurchaseSubscriptionResponse {
	return domain.PurchaseSubscriptionResponse{
		SubscriptionID:  uuid.MustParse(pb.SubscriptionId),
		ConfirmationURL: pb.ConfirmationUrl,
	}
}
