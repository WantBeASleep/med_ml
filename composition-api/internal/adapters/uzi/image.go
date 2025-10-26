package uzi

import (
	"context"

	adapter_errors "composition-api/internal/adapters/errors"
	"composition-api/internal/adapters/uzi/mappers"
	domain "composition-api/internal/domain/uzi"
	pb "composition-api/internal/generated/grpc/clients/uzi"

	"github.com/google/uuid"
)

func (a *adapter) GetImagesByUziId(ctx context.Context, id uuid.UUID) ([]domain.Image, error) {
	res, err := a.client.GetImagesByUziId(ctx, &pb.GetImagesByUziIdIn{UziId: id.String()})
	if err != nil {
		return nil, adapter_errors.HandleGRPCError(err)
	}

	return mappers.Image{}.SliceDomain(res.Images), nil
}
