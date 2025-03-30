package uzi

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/server/mappers"
	"uzi/internal/services/uzi"
)

func (h *handler) CreateUzi(ctx context.Context, in *pb.CreateUziIn) (*pb.CreateUziOut, error) {
	if _, err := uuid.Parse(in.ExternalId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "external_id is not a valid uuid: %s", err.Error())
	}

	if _, err := uuid.Parse(in.Author); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "author is not a valid uuid: %s", err.Error())
	}

	id, err := h.services.Uzi.CreateUzi(ctx, uzi.CreateUziArg{
		Projection:  mappers.UziProjectionReverseMap[in.Projection],
		ExternalID:  uuid.MustParse(in.ExternalId),
		Author:      uuid.MustParse(in.Author),
		DeviceID:    int(in.DeviceId),
		Description: in.Description,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.CreateUziOut{Id: id.String()}, nil
}
