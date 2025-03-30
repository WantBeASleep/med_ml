package uzi

import (
	"context"

	"github.com/google/uuid"

	"composition-api/internal/adapters/uzi/mappers"
	domain "composition-api/internal/domain/uzi"
	pb "composition-api/internal/generated/grpc/clients/uzi"
)

func (a *adapter) CreateUzi(ctx context.Context, in CreateUziIn) (uuid.UUID, error) {
	res, err := a.client.CreateUzi(ctx, &pb.CreateUziIn{
		Projection: in.Projection,
		ExternalId: in.ExternalID,
		DeviceId:   int64(in.DeviceID),
	})
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.MustParse(res.Id), nil
}

func (a *adapter) GetUziById(ctx context.Context, id uuid.UUID) (domain.Uzi, error) {
	res, err := a.client.GetUziById(ctx, &pb.GetUziByIdIn{Id: id.String()})
	if err != nil {
		return domain.Uzi{}, err
	}

	return mappers.Uzi(res.Uzi), nil
}

func (a *adapter) GetUzisByExternalId(ctx context.Context, id uuid.UUID) ([]domain.Uzi, error) {
	res, err := a.client.GetUzisByExternalId(ctx, &pb.GetUzisByExternalIdIn{ExternalId: id.String()})
	if err != nil {
		return nil, err
	}

	return mappers.SliceUzi(res.Uzis), nil
}

func (a *adapter) GetEchographicByUziId(ctx context.Context, id uuid.UUID) (domain.Echographic, error) {
	res, err := a.client.GetEchographicByUziId(ctx, &pb.GetEchographicByUziIdIn{UziId: id.String()})
	if err != nil {
		return domain.Echographic{}, err
	}

	return mappers.Echographic(res.Echographic), nil
}

func (a *adapter) UpdateUzi(ctx context.Context, in UpdateUziIn) (domain.Uzi, error) {
	res, err := a.client.UpdateUzi(ctx, &pb.UpdateUziIn{
		Id:         in.Id.String(),
		Projection: in.Projection,
		Checked:    in.Checked,
	})
	if err != nil {
		return domain.Uzi{}, err
	}

	return mappers.Uzi(res.Uzi), nil
}

func (a *adapter) UpdateEchographic(ctx context.Context, in domain.Echographic) (domain.Echographic, error) {
	res, err := a.client.UpdateEchographic(ctx, &pb.UpdateEchographicIn{
		Echographic: &pb.Echographic{
			Id:              in.Id.String(),
			Contors:         in.Contors,
			LeftLobeLength:  in.LeftLobeLength,
			LeftLobeWidth:   in.LeftLobeWidth,
			LeftLobeThick:   in.LeftLobeThick,
			LeftLobeVolum:   in.LeftLobeVolum,
			RightLobeLength: in.RightLobeLength,
			RightLobeWidth:  in.RightLobeWidth,
			RightLobeThick:  in.RightLobeThick,
			RightLobeVolum:  in.RightLobeVolum,
			GlandVolum:      in.GlandVolum,
			Isthmus:         in.Isthmus,
			Struct:          in.Struct,
			Echogenicity:    in.Echogenicity,
			RegionalLymph:   in.RegionalLymph,
			Vascularization: in.Vascularization,
			Location:        in.Location,
			Additional:      in.Additional,
			Conclusion:      in.Conclusion,
		},
	})
	if err != nil {
		return domain.Echographic{}, err
	}

	return mappers.Echographic(res.Echographic), nil
}
