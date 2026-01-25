package segmentation

import (
	"context"

	pb "cytology/internal/generated/grpc/service"
	"cytology/internal/server/mappers"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) GetSegmentationById(ctx context.Context, in *pb.GetSegmentationByIdIn) (*pb.GetSegmentationByIdOut, error) {
	id := int(in.Id)

	seg, err := h.services.Segmentation.GetSegmentationByID(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.GetSegmentationByIdOut{
		Segmentation: mappers.SegmentationToProto(seg),
	}, nil
}

func (h *handler) GetSegmentsByGroupId(ctx context.Context, in *pb.GetSegmentsByGroupIdIn) (*pb.GetSegmentsByGroupIdOut, error) {
	groupID := int(in.SegmentationGroupId)

	segs, err := h.services.Segmentation.GetSegmentsByGroupID(ctx, groupID)
	if err != nil {
		return &pb.GetSegmentsByGroupIdOut{Segmentations: []*pb.Segmentation{}}, nil
	}

	return &pb.GetSegmentsByGroupIdOut{
		Segmentations: mappers.SegmentationSliceToProto(segs),
	}, nil
}
