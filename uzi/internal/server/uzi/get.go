package uzi

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/repository/entity"
	"uzi/internal/server/mappers"
)

func (h *handler) GetUziById(ctx context.Context, in *pb.GetUziByIdIn) (*pb.GetUziByIdOut, error) {
	if _, err := uuid.Parse(in.Id); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id is not a valid uuid: %s", err.Error())
	}

	uzi, err := h.services.Uzi.GetUziByID(ctx, uuid.MustParse(in.Id))
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			return nil, status.Errorf(codes.NotFound, "УЗИ не найдено")
		default:
			return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
		}
	}

	out := new(pb.GetUziByIdOut)
	out.Uzi = mappers.UziFromDomain(uzi)

	return out, nil
}

func (h *handler) GetUzisByExternalId(ctx context.Context, in *pb.GetUzisByExternalIdIn) (*pb.GetUzisByExternalIdOut, error) {
	if _, err := uuid.Parse(in.ExternalId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "external_id is not a valid uuid: %s", err.Error())
	}

	uzis, err := h.services.Uzi.GetUzisByExternalID(ctx, uuid.MustParse(in.ExternalId))
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			return nil, status.Errorf(codes.NotFound, "УЗИ не найдено")
		default:
			return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
		}
	}

	out := new(pb.GetUzisByExternalIdOut)
	out.Uzis = mappers.SliceUziFromDomain(uzis)

	return out, nil
}

func (h *handler) GetUzisByAuthor(ctx context.Context, in *pb.GetUzisByAuthorIn) (*pb.GetUzisByAuthorOut, error) {
	if _, err := uuid.Parse(in.Author); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "author is not a valid uuid: %s", err.Error())
	}

	uzis, err := h.services.Uzi.GetUzisByAuthor(ctx, uuid.MustParse(in.Author))
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			return nil, status.Errorf(codes.NotFound, "УЗИ не найдено")
		default:
			return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
		}
	}

	out := new(pb.GetUzisByAuthorOut)
	out.Uzis = mappers.SliceUziFromDomain(uzis)

	return out, nil
}

func (h *handler) GetEchographicByUziId(ctx context.Context, in *pb.GetEchographicByUziIdIn) (*pb.GetEchographicByUziIdOut, error) {
	if _, err := uuid.Parse(in.UziId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "uzi_id is not a valid uuid: %s", err.Error())
	}

	echographic, err := h.services.Uzi.GetUziEchographicsByID(ctx, uuid.MustParse(in.UziId))
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			return nil, status.Errorf(codes.NotFound, "Эхографическое исследование не найдено")
		default:
			return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
		}
	}

	out := new(pb.GetEchographicByUziIdOut)
	out.Echographic = mappers.EchographicFromDomain(echographic)

	return out, nil
}
