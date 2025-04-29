package billing

import (
	"context"

	pb "composition-api/internal/generated/grpc/clients/billing"

	"github.com/google/uuid"

	domain "composition-api/internal/domain/billing"
)

type Adapter interface {
	GetTariffPlanByID(ctx context.Context, id uuid.UUID) (domain.TariffPlan, error)
	ListTariffPlans(ctx context.Context) ([]domain.TariffPlan, error)
	PurchaseSubscription(ctx context.Context, tariffPlanID, paymentProviderID, userID uuid.UUID) (domain.PurchaseSubscriptionResponse, error)
	IsUserHasActiveSubscription(ctx context.Context, userID uuid.UUID) (bool, error)
	GetUserActiveSubscription(ctx context.Context, userID uuid.UUID) (domain.Subscription, error)
	ListPaymentProviders(ctx context.Context) ([]domain.PaymentProvider, error)
	HandleYookassaWebhook(ctx context.Context, req domain.YookassaWebhookRequest) error
}

type adapter struct {
	client pb.BillingServiceClient
}

func NewAdapter(client pb.BillingServiceClient) Adapter {
	return &adapter{client: client}
}
