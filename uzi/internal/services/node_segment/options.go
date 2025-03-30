package node_segment

import "uzi/internal/domain"

// -- CreateNodesWithSegmentsOptions --

type createNodesWithSegmentsOption struct {
	newUziStatus *domain.UziStatus
}

type CreateNodesWithSegmentsOption func(*createNodesWithSegmentsOption)

var (
	WithNewUziStatus = func(status domain.UziStatus) CreateNodesWithSegmentsOption {
		return func(o *createNodesWithSegmentsOption) { o.newUziStatus = &status }
	}
)
