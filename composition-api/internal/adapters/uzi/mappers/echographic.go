package mappers

import (
	"github.com/google/uuid"

	domain "composition-api/internal/domain/uzi"
	pb "composition-api/internal/generated/grpc/clients/uzi"
)

type Echographic struct{}

func (m Echographic) Domain(pb *pb.Echographic) domain.Echographic {
	return domain.Echographic{
		Id:              uuid.MustParse(pb.Id),
		Contors:         pb.Contors,
		LeftLobeLength:  pb.LeftLobeLength,
		LeftLobeWidth:   pb.LeftLobeWidth,
		LeftLobeThick:   pb.LeftLobeThick,
		LeftLobeVolum:   pb.LeftLobeVolum,
		RightLobeLength: pb.RightLobeLength,
		RightLobeWidth:  pb.RightLobeWidth,
		RightLobeThick:  pb.RightLobeThick,
		RightLobeVolum:  pb.RightLobeVolum,
		GlandVolum:      pb.GlandVolum,
		Isthmus:         pb.Isthmus,
		Struct:          pb.Struct,
		Echogenicity:    pb.Echogenicity,
		RegionalLymph:   pb.RegionalLymph,
		Vascularization: pb.Vascularization,
		Location:        pb.Location,
		Additional:      pb.Additional,
		Conclusion:      pb.Conclusion,
	}
}

func (m Echographic) SliceDomain(pbs []*pb.Echographic) []domain.Echographic {
	return slice(pbs, m)
}
