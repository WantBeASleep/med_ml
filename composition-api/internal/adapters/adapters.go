package adapters

import (
	"google.golang.org/grpc"

	"composition-api/internal/adapters/auth"
	"composition-api/internal/adapters/med"
	"composition-api/internal/adapters/uzi"
	authPB "composition-api/internal/generated/grpc/clients/auth"
	medPB "composition-api/internal/generated/grpc/clients/med"
	uziPB "composition-api/internal/generated/grpc/clients/uzi"
)

type Adapters struct {
	Uzi  uzi.Adapter
	Auth auth.Adapter
	Med  med.Adapter
}

func NewAdapters(
	uziConn *grpc.ClientConn,
	authConn *grpc.ClientConn,
	medConn *grpc.ClientConn,
) *Adapters {
	uziClient := uziPB.NewUziSrvClient(uziConn)
	uziAdapter := uzi.NewAdapter(uziClient)

	authClient := authPB.NewAuthSrvClient(authConn)
	authAdapter := auth.NewAdapter(authClient)

	medClient := medPB.NewMedSrvClient(medConn)
	medAdapter := med.NewAdapter(medClient)

	return &Adapters{
		Uzi:  uziAdapter,
		Auth: authAdapter,
		Med:  medAdapter,
	}
}
