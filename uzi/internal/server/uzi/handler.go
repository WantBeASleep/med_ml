package uzi

import (
	"context"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/services"
)

type UziHandler interface {
	CreateUzi(ctx context.Context, req *pb.CreateUziIn) (*pb.CreateUziOut, error)

	GetUziById(ctx context.Context, req *pb.GetUziByIdIn) (*pb.GetUziByIdOut, error)
	GetUzisByExternalId(ctx context.Context, req *pb.GetUzisByExternalIdIn) (*pb.GetUzisByExternalIdOut, error)
	GetEchographicByUziId(ctx context.Context, req *pb.GetEchographicByUziIdIn) (*pb.GetEchographicByUziIdOut, error)

	UpdateUzi(ctx context.Context, req *pb.UpdateUziIn) (*pb.UpdateUziOut, error)
	UpdateEchographic(ctx context.Context, in *pb.UpdateEchographicIn) (*pb.UpdateEchographicOut, error)
}

type handler struct {
	services *services.Services
}

func New(
	services *services.Services,
) UziHandler {
	return &handler{
		services: services,
	}
}
