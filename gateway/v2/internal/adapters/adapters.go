package adapters

import (
	"google.golang.org/grpc"

	"gateway/internal/adapters/uzi"
	pb "gateway/internal/generated/grpc/clients/uzi"
)

type Adapters struct {
	Uzi uzi.Adapter
}

func NewAdapters(uziConn *grpc.ClientConn) *Adapters {
	uziClient := pb.NewUziSrvClient(uziConn)

	uziAdapter := uzi.NewAdapter(uziClient)

	return &Adapters{
		Uzi: uziAdapter,
	}
}
