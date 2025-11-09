package uzi

import (
	"context"

	adapter_errors "composition-api/internal/adapters/errors"
	"composition-api/internal/adapters/uzi/mappers"
	domain "composition-api/internal/domain/uzi"
	pb "composition-api/internal/generated/grpc/clients/uzi"

	"github.com/google/uuid"
)

var uziProjectionMap = map[domain.UziProjection]pb.UziProjection{
	domain.UziProjectionCross: pb.UziProjection_UZI_PROJECTION_CROSS,
	domain.UziProjectionLong:  pb.UziProjection_UZI_PROJECTION_LONG,
}

func (a *adapter) CreateUzi(ctx context.Context, in CreateUziIn) (uuid.UUID, error) {
	res, err := a.client.CreateUzi(ctx, &pb.CreateUziIn{
		Projection:  uziProjectionMap[in.Projection],
		ExternalId:  in.ExternalID.String(),
		Author:      in.Author.String(),
		DeviceId:    int64(in.DeviceID),
		Description: in.Description,
	})
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.MustParse(res.Id), nil
}

func (a *adapter) GetUziById(ctx context.Context, id uuid.UUID) (domain.Uzi, error) {
	res, err := a.client.GetUziById(ctx, &pb.GetUziByIdIn{Id: id.String()})
	if err != nil {
		return domain.Uzi{}, adapter_errors.HandleGRPCError(err)
	}

	return mappers.Uzi{}.Domain(res.Uzi), nil
}

func (a *adapter) GetUzisByExternalId(ctx context.Context, id uuid.UUID) ([]domain.Uzi, error) {
	res, err := a.client.GetUzisByExternalId(ctx, &pb.GetUzisByExternalIdIn{ExternalId: id.String()})
	if err != nil {
		return nil, adapter_errors.HandleGRPCError(err)
	}

	return mappers.Uzi{}.SliceDomain(res.Uzis), nil
}

func (a *adapter) GetUzisByAuthor(ctx context.Context, id uuid.UUID) ([]domain.Uzi, error) {
	res, err := a.client.GetUzisByAuthor(ctx, &pb.GetUzisByAuthorIn{Author: id.String()})
	if err != nil {
		return nil, adapter_errors.HandleGRPCError(err)
	}

	return mappers.Uzi{}.SliceDomain(res.Uzis), nil
}

func (a *adapter) GetEchographicByUziId(ctx context.Context, id uuid.UUID) (domain.Echographic, error) {
	res, err := a.client.GetEchographicByUziId(ctx, &pb.GetEchographicByUziIdIn{UziId: id.String()})
	if err != nil {
		return domain.Echographic{}, adapter_errors.HandleGRPCError(err)
	}

	return mappers.Echographic{}.Domain(res.Echographic), nil
}

func (a *adapter) UpdateUzi(ctx context.Context, in UpdateUziIn) (domain.Uzi, error) {
	res, err := a.client.UpdateUzi(ctx, &pb.UpdateUziIn{
		Id:         in.Id.String(),
		Projection: mappers.PointerFromMap(uziProjectionMap, in.Projection),
		Checked:    in.Checked,
	})
	if err != nil {
		return domain.Uzi{}, adapter_errors.HandleGRPCError(err)
	}

	return mappers.Uzi{}.Domain(res.Uzi), nil
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
		return domain.Echographic{}, adapter_errors.HandleGRPCError(err)
	}

	return mappers.Echographic{}.Domain(res.Echographic), nil
}

func (a *adapter) DeleteUzi(ctx context.Context, id uuid.UUID) error {
	_, err := a.client.DeleteUzi(ctx, &pb.DeleteUziIn{Id: id.String()})
	return err
}
