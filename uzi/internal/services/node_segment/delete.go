package node_segment

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	daoEntity "uzi/internal/repository/entity"
)

var (
	ErrChangeAiNode    = errors.New("change ai node not allowed")
	ErrChangeAiSegment = errors.New("change ai segment not allowed")
)

func (s *service) DeleteNode(ctx context.Context, id uuid.UUID) error {
	node, err := s.dao.NewNodeQuery(ctx).GetNodeByID(id)
	if err != nil {
		return fmt.Errorf("get node by id: %w", err)
	}
	if node.Ai {
		return ErrChangeAiNode
	}

	ctx, err = s.dao.BeginTx(ctx)
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

	if segment.Ai {
		return ErrChangeAiSegment
	}

	if err := segmentQuery.DeleteSegmentByID(id); err != nil {
		return fmt.Errorf("delete segment: %w", err)
	}

	_, err = segmentQuery.GetSegmentsByNodeID(segment.NodeID)
	if err != nil {
		switch {
		case errors.Is(err, daoEntity.ErrNotFound):
			// у node не осталось сегментов, удаляем
			if err := s.dao.NewNodeQuery(ctx).DeleteNodeByID(segment.NodeID); err != nil {
				return fmt.Errorf("delete node by id: %w", err)
			}

		default:
			return fmt.Errorf("get remaining segments by node_id: %w", err)
		}
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}
