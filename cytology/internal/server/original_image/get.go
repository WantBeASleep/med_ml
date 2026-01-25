package original_image

import (
	"context"

	pb "cytology/internal/generated/grpc/service"
	"cytology/internal/server/mappers"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) GetOriginalImageById(ctx context.Context, in *pb.GetOriginalImageByIdIn) (*pb.GetOriginalImageByIdOut, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id is not a valid uuid: %s", err.Error())
	}

	img, err := h.services.OriginalImage.GetOriginalImageByID(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.GetOriginalImageByIdOut{
		OriginalImage: mappers.OriginalImageToProto(img),
	}, nil
}

func (h *handler) GetOriginalImagesByCytologyId(ctx context.Context, in *pb.GetOriginalImagesByCytologyIdIn) (*pb.GetOriginalImagesByCytologyIdOut, error) {
	cytologyID, err := uuid.Parse(in.CytologyId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cytology_id is not a valid uuid: %s", err.Error())
	}

	images, err := h.services.OriginalImage.GetOriginalImagesByCytologyID(ctx, cytologyID)
	if err != nil {
		return &pb.GetOriginalImagesByCytologyIdOut{OriginalImages: []*pb.OriginalImage{}}, nil
	}

	pbImages := make([]*pb.OriginalImage, 0, len(images))
	for _, img := range images {
		pbImages = append(pbImages, mappers.OriginalImageToProto(img))
	}

	return &pb.GetOriginalImagesByCytologyIdOut{OriginalImages: pbImages}, nil
}
