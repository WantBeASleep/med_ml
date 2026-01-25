package segmentation_group

import (
	"context"

	"cytology/internal/domain"
	pb "cytology/internal/generated/grpc/service"
	"cytology/internal/server/mappers"
	"cytology/internal/services/segmentation_group"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) UpdateSegmentationGroup(ctx context.Context, in *pb.UpdateSegmentationGroupIn) (*pb.UpdateSegmentationGroupOut, error) {
	id := int(in.Id)

	var details []byte
	if in.Details != nil && *in.Details != "" {
		details = []byte(*in.Details)
	}

	var segType *domain.SegType
	if in.SegType != nil {
		st := domain.SegType(in.SegType.String())
		segType = &st
	}

	arg := segmentation_group.UpdateSegmentationGroupArg{
		Id:      id,
		SegType: segType,
		Details: details,
	}

	group, err := h.services.SegmentationGroup.UpdateSegmentationGroup(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.UpdateSegmentationGroupOut{
		SegmentationGroup: mappers.SegmentationGroupToProto(group),
	}, nil
}
