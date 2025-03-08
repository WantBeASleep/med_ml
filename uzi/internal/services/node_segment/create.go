package node_segment

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"uzi/internal/domain"
	nodeEntity "uzi/internal/repository/node/entity"
	segmentEntity "uzi/internal/repository/segment/entity"
)

const (
	avgSegmentPerNode = 3
)

func (s *service) CreateNodesWithSegments(ctx context.Context, arg []CreateNodesWithSegmentsArg) ([]CreateNodesWithSegmentsID, error) {
	result := make([]CreateNodesWithSegmentsID, 0, len(arg))
	nodes := make([]domain.Node, 0, len(arg))
	segments := make([]domain.Segment, 0, 3*len(arg))

	for _, vNode := range arg {
		nodeID := uuid.New()
		nodes = append(nodes, domain.Node{
			Id:       nodeID,
			Ai:       false,
			UziID:    vNode.Node.UziID,
			Tirads23: vNode.Node.Tirads23,
			Tirads4:  vNode.Node.Tirads4,
			Tirads5:  vNode.Node.Tirads5,
		})
		ids := CreateNodesWithSegmentsID{NodeID: nodeID}

		for _, vSegment := range vNode.Segments {
			segmentID := uuid.New()
			segments = append(segments, domain.Segment{
				Id:       segmentID,
				ImageID:  vSegment.ImageID,
				NodeID:   nodeID,
				Contor:   vSegment.Contor,
				Tirads23: vSegment.Tirads23,
				Tirads4:  vSegment.Tirads4,
				Tirads5:  vSegment.Tirads5,
			})

			ids.SegmentsID = append(ids.SegmentsID, segmentID)
		}

		result = append(result, ids)
	}

	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("begin transaction: %w", err)
	}
	defer func() { _ = s.dao.RollbackTx(ctx) }()

	nodeQuery := s.dao.NewNodeQuery(ctx)
	segmentQuery := s.dao.NewSegmentQuery(ctx)

	// вставить после переписывания репы на батчи
	if err := nodeQuery.InsertNodes(nodeEntity.Node{}.SliceFromDomain(nodes)...); err != nil {
		return nil, fmt.Errorf("insert nodes: %w", err)
	}

	if err := segmentQuery.InsertSegments(segmentEntity.Segment{}.SliceFromDomain(segments)...); err != nil {
		return nil, fmt.Errorf("insert segments: %w", err)
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return nil, fmt.Errorf("commit transaction: %w", err)
	}

	return result, nil
}
