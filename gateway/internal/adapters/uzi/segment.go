package uzi

import (
	"context"

	"github.com/google/uuid"

	"composition-api/internal/adapters/uzi/mappers"
	domain "composition-api/internal/domain/uzi"
	pb "composition-api/internal/generated/grpc/clients/uzi"
)

func (a *adapter) CreateSegment(ctx context.Context, in CreateSegmentIn) (uuid.UUID, error) {
	res, err := a.client.CreateSegment(ctx, &pb.CreateSegmentIn{
		ImageId: in.ImageID,
		NodeId:  in.NodeID,
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

	return mappers.SliceSegment(res.Segments), nil
}

func (a *adapter) UpdateSegment(ctx context.Context, in UpdateSegmentIn) (domain.Segment, error) {
	res, err := a.client.UpdateSegment(ctx, &pb.UpdateSegmentIn{
		Id:        in.Id.String(),
		Tirads_23: in.Tirads_23,
		Tirads_4:  in.Tirads_4,
		Tirads_5:  in.Tirads_5,
	})
	if err != nil {
		return domain.Segment{}, err
	}

	return mappers.Segment(res.Segment), nil
}

func (a *adapter) DeleteSegment(ctx context.Context, id uuid.UUID) error {
	_, err := a.client.DeleteSegment(ctx, &pb.DeleteSegmentIn{Id: id.String()})
	if err != nil {
		return err
	}
	return nil
}
