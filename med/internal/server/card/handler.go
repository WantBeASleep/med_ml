package card

import (
	"context"

	pb "med/internal/generated/grpc/service"
	"med/internal/services/card"

	"github.com/golang/protobuf/ptypes/empty"
)

type CardHandler interface {
	CreateCard(ctx context.Context, in *pb.CreateCardIn) (*empty.Empty, error)
	GetCard(ctx context.Context, in *pb.GetCardIn) (*pb.GetCardOut, error)
	UpdateCard(ctx context.Context, in *pb.UpdateCardIn) (*pb.UpdateCardOut, error)
}

type handler struct {
	cardSrv card.Service
}

func New(
	cardSrv card.Service,
) CardHandler {
	return &handler{
		cardSrv: cardSrv,
	}
}
