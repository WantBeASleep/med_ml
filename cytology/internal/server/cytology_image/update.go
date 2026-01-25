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

func (h *handler) UpdateCytologyImage(ctx context.Context, in *pb.UpdateCytologyImageIn) (*pb.UpdateCytologyImageOut, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id is not a valid uuid: %s", err.Error())
	}

	arg := mappers.UpdateCytologyImageArgFromProto(in, id)
	img, err := h.services.CytologyImage.UpdateCytologyImage(ctx, arg)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "cytology image not found")
		}
		if errors.Is(err, domain.ErrUnprocessableEntity) {
			return nil, status.Errorf(codes.FailedPrecondition, "Ошибка валидации данных")
		}
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.UpdateCytologyImageOut{
		CytologyImage: mappers.CytologyImageToProto(img),
	}, nil
}
