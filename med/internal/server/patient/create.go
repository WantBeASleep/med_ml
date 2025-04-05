package patient

import (
	"context"
	"time"

	"med/internal/domain"
	pb "med/internal/generated/grpc/service"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreatePatient(ctx context.Context, in *pb.CreatePatientIn) (*empty.Empty, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат id: %s", err.Error())
	}

	birthDate, err := time.Parse(time.RFC3339, in.BirthDate)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат даты рождения: %s", err.Error())
	}

	err = h.patientSrv.InsertPatient(ctx, domain.Patient{
		Id:         id,
		FullName:   in.Fullname,
		Email:      in.Email,
		Policy:     in.Policy,
		Active:     in.Active,
		Malignancy: in.Malignancy,
		BirthDate:  birthDate,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &empty.Empty{}, nil
}
