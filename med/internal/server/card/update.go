package card

import (
	"context"
	"errors"

	"med/internal/domain"
	pb "med/internal/generated/grpc/service"
	"med/internal/server/mappers"
	"med/internal/services/card"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) UpdateCard(ctx context.Context, in *pb.UpdateCardIn) (*pb.UpdateCardOut, error) {
	if _, err := uuid.Parse(in.Card.DoctorId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный ID врача: %s", err.Error())
	}
	if _, err := uuid.Parse(in.Card.PatientId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный ID пациента: %s", err.Error())
	}

	card, err := h.cardSrv.UpdateCard(
		ctx,
		uuid.MustParse(in.Card.DoctorId),
		uuid.MustParse(in.Card.PatientId),
		card.UpdateCardArg{
			Diagnosis: in.Card.Diagnosis,
		},
	)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return nil, status.Errorf(codes.NotFound, "Карта не найдена")
		case errors.Is(err, domain.ErrConflict):
			return nil, status.Errorf(codes.AlreadyExists, "Конфликт данных")
		case errors.Is(err, domain.ErrUnprocessableEntity):
			return nil, status.Errorf(codes.FailedPrecondition, "Ошибка валидации данных")
		default:
			return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
		}
	}

	return &pb.UpdateCardOut{Card: mappers.CardFromDomain(card)}, nil
}
