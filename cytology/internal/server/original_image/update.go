package original_image

import (
	"context"

	pb "cytology/internal/generated/grpc/service"
	"cytology/internal/server/mappers"
	"cytology/internal/services/original_image"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) UpdateOriginalImage(ctx context.Context, in *pb.UpdateOriginalImageIn) (*pb.UpdateOriginalImageOut, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id is not a valid uuid: %s", err.Error())
	}

	arg := original_image.UpdateOriginalImageArg{
		Id:         id,
		DelayTime:  in.DelayTime,
		ViewedFlag: in.ViewedFlag,
	}

	img, err := h.services.OriginalImage.UpdateOriginalImage(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.UpdateOriginalImageOut{
		OriginalImage: mappers.OriginalImageToProto(img),
	}, nil
}
