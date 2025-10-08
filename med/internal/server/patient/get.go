package patient

import (
	"context"
	"errors"

	pb "med/internal/generated/grpc/service"
	"med/internal/repository/entity"
	"med/internal/server/mappers"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) GetPatient(ctx context.Context, in *pb.GetPatientIn) (*pb.GetPatientOut, error) {
	patientID, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат ID пациента: %s", err.Error())
	}

	patient, err := h.patientSrv.GetPatient(ctx, patientID)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			return nil, status.Errorf(codes.NotFound, "Пациент не найден")
		default:
			return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
		}
	}

	return &pb.GetPatientOut{Patient: mappers.PatientFromDomain(patient)}, nil
}

func (h *handler) GetPatientsByDoctorID(ctx context.Context, in *pb.GetPatientsByDoctorIDIn) (*pb.GetPatientsByDoctorIDOut, error) {
	doctorID, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат ID врача: %s", err.Error())
	}

	patients, err := h.patientSrv.GetPatientsByDoctorID(ctx, doctorID)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			return nil, status.Errorf(codes.NotFound, "Пациенты не найдены")
		default:
			return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
		}
	}

	return &pb.GetPatientsByDoctorIDOut{Patients: mappers.SlicePatientFromDomain(patients)}, nil
}
