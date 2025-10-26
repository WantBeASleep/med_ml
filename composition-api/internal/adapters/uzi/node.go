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

var nodeValidationMap = map[domain.NodeValidation]pb.NodeValidation{
	domain.NodeValidationValid:   pb.NodeValidation_NODE_VALIDATION_VALID,
	domain.NodeValidationInvalid: pb.NodeValidation_NODE_VALIDATION_INVALID,
}

func (a *adapter) GetNodesByUziId(ctx context.Context, id uuid.UUID) ([]domain.Node, error) {
	res, err := a.client.GetNodesByUziId(ctx, &pb.GetNodesByUziIdIn{UziId: id.String()})
	if err != nil {
		return nil, err
	}

	return mappers.Node{}.SliceDomain(res.Nodes), nil
}

func (a *adapter) UpdateNode(ctx context.Context, in UpdateNodeIn) (domain.Node, error) {
	res, err := a.client.UpdateNode(ctx, &pb.UpdateNodeIn{
		Id:         in.Id.String(),
		Validation: mappers.PointerFromMap(nodeValidationMap, in.Validation),
		Tirads_23:  in.Tirads_23,
		Tirads_4:   in.Tirads_4,
		Tirads_5:   in.Tirads_5,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return domain.Node{}, fmt.Errorf("unknown error: %w", err)
		}

		switch st.Code() {
		case codes.NotFound:
			return domain.Node{}, adapter_errors.ErrNotFound
		default:
			return domain.Node{}, err
		}
	}

	return mappers.Node{}.Domain(res.Node), nil
}

func (a *adapter) DeleteNode(ctx context.Context, id uuid.UUID) error {
	_, err := a.client.DeleteNode(ctx, &pb.DeleteNodeIn{Id: id.String()})
	return err
}
