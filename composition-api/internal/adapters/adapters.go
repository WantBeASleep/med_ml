package adapters

import (
	billingPB "composition-api/internal/generated/grpc/clients/billing"

	"google.golang.org/grpc"

	"composition-api/internal/adapters/auth"
	"composition-api/internal/adapters/billing"
	"composition-api/internal/adapters/med"
	"composition-api/internal/adapters/uzi"
	authPB "composition-api/internal/generated/grpc/clients/auth"
	medPB "composition-api/internal/generated/grpc/clients/med"
	uziPB "composition-api/internal/generated/grpc/clients/uzi"
)

type Adapters struct {
	Uzi     uzi.Adapter
	Auth    auth.Adapter
	Med     med.Adapter
	Billing billing.Adapter
}

func NewAdapters(
	uziConn *grpc.ClientConn,
	authConn *grpc.ClientConn,
	medConn *grpc.ClientConn,
	billingConn *grpc.ClientConn,
) *Adapters {
	uziClient := uziPB.NewUziSrvClient(uziConn)
	uziAdapter := uzi.NewAdapter(uziClient)

	authClient := authPB.NewAuthSrvClient(authConn)
	authAdapter := auth.NewAdapter(authClient)

	medClient := medPB.NewMedSrvClient(medConn)
	medAdapter := med.NewAdapter(medClient)

	billingClient := billingPB.NewBillingServiceClient(billingConn)
	billingAdapter := billing.NewAdapter(billingClient)

	return &Adapters{
		Uzi:     uziAdapter,
		Auth:    authAdapter,
		Med:     medAdapter,
		Billing: billingAdapter,
	}
}
