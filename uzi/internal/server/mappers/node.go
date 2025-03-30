package mappers

import (
	"uzi/internal/domain"
	pb "uzi/internal/generated/grpc/service"
)

var nodeValidationFromDomainMap = map[domain.NodeValidation]pb.NodeValidation{
	domain.NodeValidationNull:    pb.NodeValidation_NODE_VALIDATION_NULL,
	domain.NodeValidationInvalid: pb.NodeValidation_NODE_VALIDATION_INVALID,
	domain.NodeValidationValid:   pb.NodeValidation_NODE_VALIDATION_VALID,
}

var nodeValidationToDomainMap = map[pb.NodeValidation]domain.NodeValidation{
	pb.NodeValidation_NODE_VALIDATION_NULL:    domain.NodeValidationNull,
	pb.NodeValidation_NODE_VALIDATION_INVALID: domain.NodeValidationInvalid,
	pb.NodeValidation_NODE_VALIDATION_VALID:   domain.NodeValidationValid,
}

func NodeValidationFromDomain(validation *domain.NodeValidation) *pb.NodeValidation {
	if validation == nil {
		return nil
	}
	v := nodeValidationFromDomainMap[*validation]
	return &v
}

func NodeValidationToDomain(validation *pb.NodeValidation) *domain.NodeValidation {
	if validation == nil {
		return nil
	}
	v := nodeValidationToDomainMap[*validation]
	return &v
}

func NodeFromDomain(domain domain.Node) *pb.Node {
	return &pb.Node{
		Id:         domain.Id.String(),
		Ai:         domain.Ai,
		Validation: NodeValidationFromDomain(domain.Validation),
		UziId:      domain.UziID.String(),
		Tirads_23:  domain.Tirads23,
		Tirads_4:   domain.Tirads4,
		Tirads_5:   domain.Tirads5,
	}
}

func SliceNodeFromDomain(domains []domain.Node) []*pb.Node {
	pbs := make([]*pb.Node, 0, len(domains))
	for _, d := range domains {
		pbs = append(pbs, NodeFromDomain(d))
	}
	return pbs
}
