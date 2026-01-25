package segmentation_group

import (
	"context"

	pb "cytology/internal/generated/grpc/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *handler) DeleteSegmentationGroup(ctx context.Context, in *pb.DeleteSegmentationGroupIn) (*emptypb.Empty, error) {
	id := int(in.Id)

	err := h.services.SegmentationGroup.DeleteSegmentationGroup(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &emptypb.Empty{}, nil
}
