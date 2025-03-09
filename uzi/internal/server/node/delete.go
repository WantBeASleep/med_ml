package node

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "uzi/internal/generated/grpc/service"
)

func (h *handler) DeleteNode(ctx context.Context, in *pb.DeleteNodeIn) (*empty.Empty, error) {
	if _, err := uuid.Parse(in.Id); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id is not a valid uuid: %s", err.Error())
	}

	if err := h.services.NodeSegment.DeleteNode(ctx, uuid.MustParse(in.Id)); err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}
	return &empty.Empty{}, nil
}
