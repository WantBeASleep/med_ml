package card

import (
	"context"

	"med/internal/domain"
	pb "med/internal/generated/grpc/service"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateCard(ctx context.Context, in *pb.CreateCardIn) (*empty.Empty, error) {
	if _, err := uuid.Parse(in.Card.DoctorId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный ID врача: %s", err.Error())
	}
	if _, err := uuid.Parse(in.Card.PatientId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный ID пациента: %s", err.Error())
	}

	if err := h.cardSrv.CreateCard(ctx, domain.Card{
		DoctorID:  uuid.MustParse(in.Card.DoctorId),
		PatientID: uuid.MustParse(in.Card.PatientId),
		Diagnosis: in.Card.Diagnosis,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &empty.Empty{}, nil
}
