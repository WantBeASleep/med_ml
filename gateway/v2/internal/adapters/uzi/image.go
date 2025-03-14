package uzi

import (
	"context"

	"github.com/google/uuid"

	"gateway/internal/adapters/uzi/mappers"
	domain "gateway/internal/domain/uzi"
	pb "gateway/internal/generated/grpc/clients/uzi"
)

func (a *adapter) GetImagesByUziId(ctx context.Context, id uuid.UUID) ([]domain.Image, error) {
	res, err := a.client.GetImagesByUziId(ctx, &pb.GetImagesByUziIdIn{UziId: id.String()})
	if err != nil {
		return nil, err
	}

	return mappers.SliceImage(res.Images), nil
}
