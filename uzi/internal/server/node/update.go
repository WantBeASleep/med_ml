package node

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/server/mappers"
	"uzi/internal/services/node"
)

func (h *handler) UpdateNode(ctx context.Context, in *pb.UpdateNodeIn) (*pb.UpdateNodeOut, error) {
	if _, err := uuid.Parse(in.Id); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id is not a valid uuid: %s", err.Error())
	}

	node, err := h.services.Node.UpdateNode(
		ctx,
		node.UpdateNodeArg{
			Id:         uuid.MustParse(in.Id),
			Validation: mappers.NodeValidationToDomain(in.Validation),
			Tirads23:   in.Tirads_23,
			Tirads4:    in.Tirads_4,
			Tirads5:    in.Tirads_5,
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	out := new(pb.UpdateNodeOut)
	out.Node = mappers.NodeFromDomain(node)

	return out, nil
}
