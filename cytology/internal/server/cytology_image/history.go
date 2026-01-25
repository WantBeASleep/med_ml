package cytology_image

import (
	"context"

	pb "cytology/internal/generated/grpc/service"
	"cytology/internal/server/mappers"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) GetCytologyImageHistory(ctx context.Context, in *pb.GetCytologyImageHistoryIn) (*pb.GetCytologyImageHistoryOut, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id is not a valid uuid: %s", err.Error())
	}

	images, err := h.services.CytologyImage.GetCytologyImageHistory(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	pbImages := make([]*pb.CytologyImage, 0, len(images))
	for _, img := range images {
		pbImages = append(pbImages, mappers.CytologyImageToProto(img))
	}

	return &pb.GetCytologyImageHistoryOut{CytologyImages: pbImages}, nil
}
