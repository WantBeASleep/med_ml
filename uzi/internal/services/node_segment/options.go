package node_segment

import "uzi/internal/domain"

// -- CreateNodesWithSegmentsOptions --

type createNodesWithSegmentsOption struct {
	newUziStatus       *domain.UziStatus
	setNodesValidation *domain.NodeValidation
}

type CreateNodesWithSegmentsOption func(*createNodesWithSegmentsOption)

var (
	WithNewUziStatus = func(status domain.UziStatus) CreateNodesWithSegmentsOption {
		return func(o *createNodesWithSegmentsOption) { o.newUziStatus = &status }
	}

	WithSetNodesValidation = func(validation domain.NodeValidation) CreateNodesWithSegmentsOption {
		return func(o *createNodesWithSegmentsOption) { o.setNodesValidation = &validation }
	}
)
