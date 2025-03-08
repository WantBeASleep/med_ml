package mappers

import (
	"uzi/internal/domain"
	pb "uzi/internal/generated/grpc/service"
)

func EchographicFromDomain(domain domain.Echographic) *pb.Echographic {
	return &pb.Echographic{
		Id:              domain.Id.String(),
		Contors:         domain.Contors,
		LeftLobeLength:  domain.LeftLobeLength,
		LeftLobeWidth:   domain.LeftLobeWidth,
		LeftLobeThick:   domain.LeftLobeThick,
		LeftLobeVolum:   domain.LeftLobeVolum,
		RightLobeLength: domain.RightLobeLength,
		RightLobeWidth:  domain.RightLobeWidth,
		RightLobeThick:  domain.RightLobeThick,
		RightLobeVolum:  domain.RightLobeVolum,
		GlandVolum:      domain.GlandVolum,
		Isthmus:         domain.Isthmus,
		Struct:          domain.Struct,
		Echogenicity:    domain.Echogenicity,
		RegionalLymph:   domain.RegionalLymph,
		Vascularization: domain.Vascularization,
		Location:        domain.Location,
		Additional:      domain.Additional,
		Conclusion:      domain.Conclusion,
	}
}

func SliceEchographicFromDomain(domains []domain.Echographic) []*pb.Echographic {
	pbs := make([]*pb.Echographic, 0, len(domains))
	for _, d := range domains {
		pbs = append(pbs, EchographicFromDomain(d))
	}
	return pbs
}
