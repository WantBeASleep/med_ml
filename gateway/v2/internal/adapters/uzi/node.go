package uzi

import (
	"context"

	"github.com/google/uuid"

	"gateway/internal/adapters/uzi/mappers"
	domain "gateway/internal/domain/uzi"
	pb "gateway/internal/generated/grpc/clients/uzi"
)

func (a *adapter) GetNodesByUziId(ctx context.Context, id uuid.UUID) ([]domain.Node, error) {
	res, err := a.client.GetNodesByUziId(ctx, &pb.GetNodesByUziIdIn{UziId: id.String()})
	if err != nil {
		return nil, err
	}

	return mappers.SliceNode(res.Nodes), nil
}

func (a *adapter) UpdateNode(ctx context.Context, in UpdateNodeIn) (domain.Node, error) {
	res, err := a.client.UpdateNode(ctx, &pb.UpdateNodeIn{
		Id: in.Id.String(),
	})
	if err != nil {
		return domain.Node{}, err
	}

	return mappers.Node(res.Node), nil
}

func (a *adapter) DeleteNode(ctx context.Context, id uuid.UUID) error {
	_, err := a.client.DeleteNode(ctx, &pb.DeleteNodeIn{Id: id.String()})
	if err != nil {
		return err
	}
	return nil
}
