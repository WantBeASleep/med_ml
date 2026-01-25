package cytology_image

import (
	"context"
	"errors"

	"cytology/internal/domain"
	pb "cytology/internal/generated/grpc/service"
	"cytology/internal/server/mappers"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CopyCytologyImage(ctx context.Context, in *pb.CopyCytologyImageIn) (*pb.CopyCytologyImageOut, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id is not a valid uuid: %s", err.Error())
	}

	img, err := h.services.CytologyImage.CopyCytologyImage(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrBadRequest) {
			return nil, status.Errorf(codes.FailedPrecondition, "can only copy last version")
		}
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.CopyCytologyImageOut{
		CytologyImage: mappers.CytologyImageToProto(img),
	}, nil
}
