package uzi

import (
	"context"
	"fmt"

	adapter_errors "composition-api/internal/adapters/errors"
	"composition-api/internal/adapters/uzi/mappers"
	domain "composition-api/internal/domain/uzi"
	pb "composition-api/internal/generated/grpc/clients/uzi"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *adapter) GetImagesByUziId(ctx context.Context, id uuid.UUID) ([]domain.Image, error) {
	res, err := a.client.GetImagesByUziId(ctx, &pb.GetImagesByUziIdIn{UziId: id.String()})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, fmt.Errorf("unknown error: %w", err)
		}

		switch st.Code() {
		case codes.NotFound:
			return nil, adapter_errors.ErrNotFound
		default:
			return nil, err
		}
	}

	return mappers.Image{}.SliceDomain(res.Images), nil
}
