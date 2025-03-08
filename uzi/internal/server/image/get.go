package image

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/server/mappers"
)

func (h *handler) GetImagesByUziId(ctx context.Context, in *pb.GetImagesByUziIdIn) (*pb.GetImagesByUziIdOut, error) {
	if _, err := uuid.Parse(in.UziId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "uzi_id is not a valid uuid: %s", err.Error())
	}

	images, err := h.services.Image.GetImagesByUziID(ctx, uuid.MustParse(in.UziId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	out := new(pb.GetImagesByUziIdOut)
	out.Images = mappers.SliceImageFromDomain(images)

	return out, nil
}
