package segmentation

import (
	"context"

	pb "cytology/internal/generated/grpc/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *handler) DeleteSegmentation(ctx context.Context, in *pb.DeleteSegmentationIn) (*emptypb.Empty, error) {
	id := int(in.Id)

	err := h.services.Segmentation.DeleteSegmentation(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &emptypb.Empty{}, nil
}
