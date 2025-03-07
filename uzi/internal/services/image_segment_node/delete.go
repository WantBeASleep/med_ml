package image_segment_node

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *service) DeleteNode(ctx context.Context, id uuid.UUID) error {
	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer func() { _ = s.dao.RollbackTx(ctx) }()

	if err := s.dao.NewSegmentQuery(ctx).DeleteSegmentsByUziID(id); err != nil {
		return fmt.Errorf("delete node segments: %w", err)
	}

	if err := s.dao.NewNodeQuery(ctx).DeleteNodeByID(id); err != nil {
		return fmt.Errorf("delete node: %w", err)
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (s *service) DeleteSegment(ctx context.Context, id uuid.UUID) error {
	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer func() { _ = s.dao.RollbackTx(ctx) }()

	segmentQuery := s.dao.NewSegmentQuery(ctx)

	segment, err := segmentQuery.GetSegmentByID(id)
	if err != nil {
		return fmt.Errorf("get segment by id: %w", err)
	}

	if err := segmentQuery.DeleteSegmentByID(id); err != nil {
		return fmt.Errorf("delete segment: %w", err)
	}

	remainingSegments, err := segmentQuery.GetSegmentsByNodeID(segment.NodeID)
	if err != nil {
		return fmt.Errorf("get segment by node_id: %w", err)
	}

	// у node не осталось сегментов, удаляем
	if len(remainingSegments) == 0 {
		if err := s.dao.NewNodeQuery(ctx).DeleteNodeByID(segment.NodeID); err != nil {
			return fmt.Errorf("delete node by id: %w", err)
		}
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}
