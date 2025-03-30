package patient

import (
	"context"
	"time"

	"med/internal/domain"
	pb "med/internal/generated/grpc/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreatePatient(ctx context.Context, in *pb.CreatePatientIn) (*pb.CreatePatientOut, error) {
	birthDate, err := time.Parse(time.RFC3339, in.BirthDate)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат даты рождения: %s", err.Error())
	}

	id, err := h.patientSrv.CreatePatient(ctx, domain.Patient{
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

	return &pb.CreatePatientOut{Id: id.String()}, nil
}
