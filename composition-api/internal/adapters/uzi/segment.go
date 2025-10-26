package uzi

import (
	"context"

	adapter_errors "composition-api/internal/adapters/errors"
	"composition-api/internal/adapters/uzi/mappers"
	domain "composition-api/internal/domain/uzi"
	pb "composition-api/internal/generated/grpc/clients/uzi"

	"github.com/google/uuid"
)

func (a *adapter) CreateSegment(ctx context.Context, in CreateSegmentIn) (uuid.UUID, error) {
	res, err := a.client.CreateSegment(ctx, &pb.CreateSegmentIn{
		ImageId:   in.ImageID.String(),
		NodeId:    in.NodeID.String(),
		Contor:    in.Contor,
		Tirads_23: in.Tirads_23,
		Tirads_4:  in.Tirads_4,
		Tirads_5:  in.Tirads_5,
	})
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.MustParse(res.Id), nil
}

func (a *adapter) GetSegmentsByNodeId(ctx context.Context, id uuid.UUID) ([]domain.Segment, error) {
	res, err := a.client.GetSegmentsByNodeId(ctx, &pb.GetSegmentsByNodeIdIn{NodeId: id.String()})
	if err != nil {
		return nil, err
	}

	return mappers.Segment{}.SliceDomain(res.Segments), nil
}

func (a *adapter) UpdateSegment(ctx context.Context, in UpdateSegmentIn) (domain.Segment, error) {
	res, err := a.client.UpdateSegment(ctx, &pb.UpdateSegmentIn{
		Id:        in.Id.String(),
		Contor:    in.Contor,
		Tirads_23: in.Tirads_23,
		Tirads_4:  in.Tirads_4,
		Tirads_5:  in.Tirads_5,
	})
	if err != nil {
		return domain.Segment{}, adapter_errors.HandleGRPCError(err)
	}

	return mappers.Segment{}.Domain(res.Segment), nil
}

func (a *adapter) DeleteSegment(ctx context.Context, id uuid.UUID) error {
	_, err := a.client.DeleteSegment(ctx, &pb.DeleteSegmentIn{Id: id.String()})
	return err
}
