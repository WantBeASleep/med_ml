package uzi

import (
	"context"

	"github.com/google/uuid"

	pb "uzi/internal/generated/grpc/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *handler) DeleteUzi(ctx context.Context, req *pb.DeleteUziIn) (*emptypb.Empty, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid uzi id: %v", err)
	}

	err = h.services.Uzi.DeleteUzi(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete uzi: %v", err)
	}

	return &emptypb.Empty{}, nil
}
