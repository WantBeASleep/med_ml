package mappers

import (
	"github.com/AlekSi/pointer"
	"github.com/google/uuid"

	domain "composition-api/internal/domain/uzi"
	pb "composition-api/internal/generated/grpc/clients/uzi"
)

var nodeValidationMap = map[pb.NodeValidation]domain.NodeValidation{
	pb.NodeValidation_NODE_VALIDATION_NULL:    domain.NodeValidationNull,
	pb.NodeValidation_NODE_VALIDATION_VALID:   domain.NodeValidationValid,
	pb.NodeValidation_NODE_VALIDATION_INVALID: domain.NodeValidationInvalid,
}

func nodeValidation(pb *pb.NodeValidation) *domain.NodeValidation {
	if pb == nil {
		return nil
	}
	return pointer.To(nodeValidationMap[*pb])
}

type Node struct{}

func (m Node) Domain(pb *pb.Node) domain.Node {
	return domain.Node{
		Id:          uuid.MustParse(pb.Id),
		Ai:          pb.Ai,
		UziID:       uuid.MustParse(pb.UziId),
		Validation:  nodeValidation(pb.Validation),
		Tirads23:    pb.Tirads_23,
		Tirads4:     pb.Tirads_4,
		Tirads5:     pb.Tirads_5,
		Description: pb.Description,
	}
}

func (m Node) SliceDomain(pbs []*pb.Node) []domain.Node {
	return slice(pbs, m)
}
