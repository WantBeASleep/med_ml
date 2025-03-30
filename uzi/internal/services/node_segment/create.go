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

func (s *service) createNodesWithSegments(
	ctx context.Context,
	uziID uuid.UUID,
	ai bool,
	arg []CreateNodesWithSegmentsArg,
	opts ...CreateNodesWithSegmentsOption,
) ([]CreateNodesWithSegmentsID, error) {
	opt := &createNodesWithSegmentsOption{}
	for _, o := range opts {
		o(opt)
	}

	nodes, segments, ids := s.createDomainNodeSegmentsFromArgs(uziID, ai, arg)

	if opt.setNodesValidation != nil {
		for i := range nodes {
			nodes[i].Validation = opt.setNodesValidation
		}
	}

	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("begin transaction: %w", err)
	}
	defer func() { _ = s.dao.RollbackTx(ctx) }()

	nodeQuery := s.dao.NewNodeQuery(ctx)
	segmentQuery := s.dao.NewSegmentQuery(ctx)

	if err := nodeQuery.InsertNodes(nodeEntity.Node{}.SliceFromDomain(nodes)...); err != nil {
		return nil, fmt.Errorf("insert nodes: %w", err)
	}

	if err := segmentQuery.InsertSegments(segmentEntity.Segment{}.SliceFromDomain(segments)...); err != nil {
		return nil, fmt.Errorf("insert segments: %w", err)
	}

	if opt.newUziStatus != nil {
		if err := s.dao.NewUziQuery(ctx).UpdateUziStatus(uziID, opt.newUziStatus.String()); err != nil {
			return nil, fmt.Errorf("update uzi status: %w", err)
		}
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return nil, fmt.Errorf("commit transaction: %w", err)
	}

	return ids, nil
}

func (s *service) SaveProcessedNodesWithSegments(
	ctx context.Context,
	uziID uuid.UUID,
	arg []CreateNodesWithSegmentsArg,
) error {
	_, err := s.createNodesWithSegments(
		ctx,
		uziID,
		true,
		arg,
		WithNewUziStatus(domain.UziStatusCompleted),
		WithSetNodesValidation(domain.NodeValidationNull),
	)

	return err
}

func (s *service) CreateManualNodesWithSegments(
	ctx context.Context,
	uziID uuid.UUID,
	arg []CreateNodesWithSegmentsArg,
) ([]CreateNodesWithSegmentsID, error) {
	return s.createNodesWithSegments(ctx, uziID, false, arg)
}

func (s *service) createDomainNodeSegmentsFromArgs(
	uziID uuid.UUID,
	ai bool,
	arg []CreateNodesWithSegmentsArg,
) (
	[]domain.Node,
	[]domain.Segment,
	[]CreateNodesWithSegmentsID,
) {
	ids := make([]CreateNodesWithSegmentsID, 0, len(arg))
	nodes := make([]domain.Node, 0, len(arg))
	segments := make([]domain.Segment, 0, avgSegmentPerNode*len(arg))

	for _, NodeAndSeg := range arg {
		nodeID := uuid.New()
		nodes = append(nodes, domain.Node{
			Id:       nodeID,
			Ai:       ai,
			UziID:    uziID,
			Tirads23: NodeAndSeg.Node.Tirads23,
			Tirads4:  NodeAndSeg.Node.Tirads4,
			Tirads5:  NodeAndSeg.Node.Tirads5,
		})

		id := CreateNodesWithSegmentsID{NodeID: nodeID}

		for _, segment := range NodeAndSeg.Segments {
			segmentID := uuid.New()
			segments = append(segments, domain.Segment{
				Id:       segmentID,
				ImageID:  segment.ImageID,
				NodeID:   nodeID,
				Contor:   segment.Contor,
				Ai:       ai,
				Tirads23: segment.Tirads23,
				Tirads4:  segment.Tirads4,
				Tirads5:  segment.Tirads5,
			})

			id.SegmentsID = append(id.SegmentsID, segmentID)
		}

		ids = append(ids, id)
	}

	return nodes, segments, ids
}
