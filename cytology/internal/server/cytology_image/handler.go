package cytology_image

import (
	"context"

	pb "cytology/internal/generated/grpc/service"
	"cytology/internal/services"

	"google.golang.org/protobuf/types/known/emptypb"
)

type CytologyImageHandler interface {
	CreateCytologyImage(ctx context.Context, req *pb.CreateCytologyImageIn) (*pb.CreateCytologyImageOut, error)
	GetCytologyImageById(ctx context.Context, req *pb.GetCytologyImageByIdIn) (*pb.GetCytologyImageByIdOut, error)
	GetCytologyImagesByExternalId(ctx context.Context, req *pb.GetCytologyImagesByExternalIdIn) (*pb.GetCytologyImagesByExternalIdOut, error)
	GetCytologyImagesByDoctorIdAndPatientId(ctx context.Context, req *pb.GetCytologyImagesByDoctorIdAndPatientIdIn) (*pb.GetCytologyImagesByDoctorIdAndPatientIdOut, error)
	UpdateCytologyImage(ctx context.Context, req *pb.UpdateCytologyImageIn) (*pb.UpdateCytologyImageOut, error)
	DeleteCytologyImage(ctx context.Context, req *pb.DeleteCytologyImageIn) (*emptypb.Empty, error)
	CopyCytologyImage(ctx context.Context, req *pb.CopyCytologyImageIn) (*pb.CopyCytologyImageOut, error)
	GetCytologyImageHistory(ctx context.Context, req *pb.GetCytologyImageHistoryIn) (*pb.GetCytologyImageHistoryOut, error)
}

type handler struct {
	services *services.Services
}

func New(services *services.Services) CytologyImageHandler {
	return &handler{
		services: services,
	}
}
