package segmentation

import (
	"context"

	"cytology/internal/domain"
	pb "cytology/internal/generated/grpc/service"
	"cytology/internal/services/segmentation"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateSegmentation(ctx context.Context, in *pb.CreateSegmentationIn) (*pb.CreateSegmentationOut, error) {
	groupID := int(in.SegmentationGroupId)

	points := make([]domain.SegmentationPoint, 0, len(in.Points))
	for _, p := range in.Points {
		points = append(points, domain.SegmentationPoint{
			X: int(p.X),
			Y: int(p.Y),
		})
	}

	arg := segmentation.CreateSegmentationArg{
		SegmentationGroupID: groupID,
		Points:              points,
	}

	id, err := h.services.Segmentation.CreateSegmentation(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.CreateSegmentationOut{Id: int32(id)}, nil
}
