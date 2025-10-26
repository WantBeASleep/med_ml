package card

import (
	"context"
	"errors"
	"log/slog"

	pb "med/internal/generated/grpc/service"
	"med/internal/repository/entity"
	"med/internal/server/mappers"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) GetCard(ctx context.Context, in *pb.GetCardIn) (*pb.GetCardOut, error) {
	slog.Info("GetCard called", "doctorId", in.DoctorId, "patientId", in.PatientId)

	if _, err := uuid.Parse(in.DoctorId); err != nil {
		slog.Error("Invalid doctor ID", "doctorId", in.DoctorId, "err", err)
		return nil, status.Errorf(codes.InvalidArgument, "Неверный ID врача: %s", err.Error())
	}
	if _, err := uuid.Parse(in.PatientId); err != nil {
		slog.Error("Invalid patient ID", "patientId", in.PatientId, "err", err)
		return nil, status.Errorf(codes.InvalidArgument, "Неверный ID пациента: %s", err.Error())
	}

	doctorUUID := uuid.MustParse(in.DoctorId)
	patientUUID := uuid.MustParse(in.PatientId)

	slog.Info("Calling service GetCard", "doctorId", doctorUUID, "patientId", patientUUID)
	card, err := h.cardSrv.GetCard(ctx, doctorUUID, patientUUID)
	if err != nil {
		slog.Error("Error getting card", "doctorId", doctorUUID, "patientId", patientUUID, "err", err)
		switch {
		case errors.Is(err, entity.ErrNotFound):
			slog.Warn("Card not found", "doctorId", doctorUUID, "patientId", patientUUID)
			return nil, status.Errorf(codes.NotFound, "Карта не найдена")
		default:
			return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
		}
	}

	slog.Info("Card retrieved successfully", "doctorId", doctorUUID, "patientId", patientUUID, "diagnosis", card.Diagnosis)
	return &pb.GetCardOut{Card: mappers.CardFromDomain(card)}, nil
}
