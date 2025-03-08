package node_segment

import "uzi/internal/domain"

type createNodesWithSegmentsOption struct {
	newUziStatus *domain.UziStatus
}

type CreateNodesWithSegmentsOption func(*createNodesWithSegmentsOption)

func WithNewUziStatus(status domain.UziStatus) CreateNodesWithSegmentsOption {
	return func(o *createNodesWithSegmentsOption) { o.newUziStatus = &status }
}
