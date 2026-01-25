package adapters

import (
	billingPB "composition-api/internal/generated/grpc/clients/billing"

	"google.golang.org/grpc"

	"composition-api/internal/adapters/auth"
	"composition-api/internal/adapters/billing"
	"composition-api/internal/adapters/cytology"
	"composition-api/internal/adapters/med"
	"composition-api/internal/adapters/tiler"
	"composition-api/internal/adapters/uzi"
	authPB "composition-api/internal/generated/grpc/clients/auth"
	cytologyPB "composition-api/internal/generated/grpc/clients/cytology"
	medPB "composition-api/internal/generated/grpc/clients/med"
	uziPB "composition-api/internal/generated/grpc/clients/uzi"
)

type Adapters struct {
	Uzi      uzi.Adapter
	Auth     auth.Adapter
	Med      med.Adapter
	Billing  billing.Adapter
	Cytology cytology.Adapter
	Tiler    tiler.Client
}

func NewAdapters(
	uziConn *grpc.ClientConn,
	authConn *grpc.ClientConn,
	medConn *grpc.ClientConn,
	billingConn *grpc.ClientConn,
	cytologyConn *grpc.ClientConn,
	tilerURL string,
) *Adapters {
	uziClient := uziPB.NewUziSrvClient(uziConn)
	uziAdapter := uzi.NewAdapter(uziClient)

	authClient := authPB.NewAuthSrvClient(authConn)
	authAdapter := auth.NewAdapter(authClient)

	medClient := medPB.NewMedSrvClient(medConn)
	medAdapter := med.NewAdapter(medClient)

	billingClient := billingPB.NewBillingServiceClient(billingConn)
	billingAdapter := billing.NewAdapter(billingClient)

	cytologyClient := cytologyPB.NewCytologySrvClient(cytologyConn)
	cytologyAdapter := cytology.NewAdapter(cytologyClient)

	tilerClient := tiler.NewClient(tilerURL)

	return &Adapters{
		Uzi:      uziAdapter,
		Auth:     authAdapter,
		Med:      medAdapter,
		Billing:  billingAdapter,
		Cytology: cytologyAdapter,
		Tiler:    tilerClient,
	}
}
