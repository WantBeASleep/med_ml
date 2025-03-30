package card

import (
	"context"

	pb "med/internal/generated/grpc/service"
	"med/internal/server/mappers"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) GetCard(ctx context.Context, in *pb.GetCardIn) (*pb.GetCardOut, error) {
	if _, err := uuid.Parse(in.DoctorId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный ID врача: %s", err.Error())
	}
	if _, err := uuid.Parse(in.PatientId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный ID пациента: %s", err.Error())
	}

	card, err := h.cardSrv.GetCard(ctx, uuid.MustParse(in.DoctorId), uuid.MustParse(in.PatientId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.GetCardOut{Card: mappers.CardFromDomain(card)}, nil
}
