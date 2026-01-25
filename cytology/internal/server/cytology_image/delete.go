package cytology_image

import (
	"context"
	"errors"

	"cytology/internal/domain"
	pb "cytology/internal/generated/grpc/service"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *handler) DeleteCytologyImage(ctx context.Context, in *pb.DeleteCytologyImageIn) (*emptypb.Empty, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id is not a valid uuid: %s", err.Error())
	}

	err = h.services.CytologyImage.DeleteCytologyImage(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "cytology image not found")
		}
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &emptypb.Empty{}, nil
}
