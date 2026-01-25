package card

import (
	"context"
	"errors"

	"med/internal/domain"
	pb "med/internal/generated/grpc/service"
	"med/internal/server/mappers"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateCard(ctx context.Context, in *pb.CreateCardIn) (*pb.CreateCardOut, error) {
	if _, err := uuid.Parse(in.Card.DoctorId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный ID врача: %s", err.Error())
	}
	if _, err := uuid.Parse(in.Card.PatientId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный ID пациента: %s", err.Error())
	}

	createdCard, err := h.cardSrv.CreateCard(ctx, domain.Card{
		DoctorID:  uuid.MustParse(in.Card.DoctorId),
		PatientID: uuid.MustParse(in.Card.PatientId),
		Diagnosis: in.Card.Diagnosis,
	})
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return nil, status.Errorf(codes.NotFound, "Пациент не найден")
		case errors.Is(err, domain.ErrBadRequest):
			return nil, status.Errorf(codes.InvalidArgument, "Неверный формат ОМС пациента")
		case errors.Is(err, domain.ErrConflict):
			return nil, status.Errorf(codes.AlreadyExists, "Конфликт данных")
		case errors.Is(err, domain.ErrUnprocessableEntity):
			return nil, status.Errorf(codes.FailedPrecondition, "Ошибка валидации данных")
		default:
			return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
		}
	}

	return &pb.CreateCardOut{
		Card: mappers.CardFromDomain(createdCard),
	}, nil
}
