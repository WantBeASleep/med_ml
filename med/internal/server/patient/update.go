package patient

import (
	"context"
	"errors"
	"time"

	pb "med/internal/generated/grpc/service"
	"med/internal/repository/entity"
	"med/internal/server/mappers"
	"med/internal/services/patient"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) UpdatePatient(ctx context.Context, in *pb.UpdatePatientIn) (*pb.UpdatePatientOut, error) {
	patientID, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат ID пациента: %s", err.Error())
	}

	var lastUziDate *time.Time
	if in.LastUziDate != nil {
		lastUziDateParsed, err := time.Parse(time.RFC3339, *in.LastUziDate)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Неверный формат даты последнего УЗИ: %s", err.Error())
		}
		lastUziDate = &lastUziDateParsed
	}

	patient, err := h.patientSrv.UpdatePatient(
		ctx,
		patientID,
		patient.UpdatePatient{
			Active:      in.Active,
			Malignancy:  in.Malignancy,
			LastUziDate: lastUziDate,
		},
	)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			return nil, status.Errorf(codes.NotFound, "Пациент не найден")
		default:
			return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
		}
	}

	return &pb.UpdatePatientOut{Patient: mappers.PatientFromDomain(patient)}, nil
}
