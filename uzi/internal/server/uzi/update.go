package uzi

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"uzi/internal/domain"

	"github.com/AlekSi/pointer"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/server/mappers"
	"uzi/internal/services/uzi"
)

func (h *handler) UpdateUzi(ctx context.Context, in *pb.UpdateUziIn) (*pb.UpdateUziOut, error) {
	if _, err := uuid.Parse(in.Id); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id is not a valid uuid: %s", err.Error())
	}

	var projection *domain.UziProjection
	if in.Projection != nil {
		projection = pointer.To(mappers.UziProjectionReverseMap[*in.Projection])
	}

	uzi, err := h.services.Uzi.UpdateUzi(ctx, uzi.UpdateUziArg{
		Id:         uuid.MustParse(in.Id),
		Projection: projection,
		Checked:    in.Checked,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	out := new(pb.UpdateUziOut)
	out.Uzi = mappers.UziFromDomain(uzi)

	return out, nil
}

func (h *handler) UpdateEchographic(ctx context.Context, in *pb.UpdateEchographicIn) (*pb.UpdateEchographicOut, error) {
	echographic, err := h.services.Uzi.UpdateEchographic(
		ctx,
		uzi.UpdateEchographicArg{
			Id:              uuid.MustParse(in.Echographic.Id),
			Contors:         in.Echographic.Contors,
			LeftLobeLength:  in.Echographic.LeftLobeLength,
			LeftLobeWidth:   in.Echographic.LeftLobeWidth,
			LeftLobeThick:   in.Echographic.LeftLobeThick,
			LeftLobeVolum:   in.Echographic.LeftLobeVolum,
			RightLobeLength: in.Echographic.RightLobeLength,
			RightLobeWidth:  in.Echographic.RightLobeWidth,
			RightLobeThick:  in.Echographic.RightLobeThick,
			RightLobeVolum:  in.Echographic.RightLobeVolum,
			GlandVolum:      in.Echographic.GlandVolum,
			Isthmus:         in.Echographic.Isthmus,
			Struct:          in.Echographic.Struct,
			Echogenicity:    in.Echographic.Echogenicity,
			RegionalLymph:   in.Echographic.RegionalLymph,
			Vascularization: in.Echographic.Vascularization,
			Location:        in.Echographic.Location,
			Additional:      in.Echographic.Additional,
			Conclusion:      in.Echographic.Conclusion,
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	out := new(pb.UpdateEchographicOut)
	out.Echographic = mappers.EchographicFromDomain(echographic)

	return out, nil
}
