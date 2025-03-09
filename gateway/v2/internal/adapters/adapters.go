package adapters

import (
	"gateway/internal/adapters/uzi"
	pb "gateway/internal/generated/grpc/clients/uzi"
)

type Adapters struct {
	Uzi uzi.Adapter
}

func NewAdapters(client pb.UziSrvClient) *Adapters {
	Uzi := uzi.NewAdapter(client)

	return &Adapters{
		Uzi: Uzi,
	}
}
