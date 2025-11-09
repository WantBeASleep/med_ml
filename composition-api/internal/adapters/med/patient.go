package med

import (
	"context"
	"log/slog"
	"time"

	adapter_errors "composition-api/internal/adapters/errors"
	"composition-api/internal/adapters/med/mappers"
	domain "composition-api/internal/domain/med"
	pb "composition-api/internal/generated/grpc/clients/med"

	"github.com/google/uuid"
)

func (a *adapter) CreatePatient(ctx context.Context, arg CreatePatientArg) error {
	_, err := a.client.CreatePatient(ctx, &pb.CreatePatientIn{
		Id:         arg.Id.String(),
		Fullname:   arg.FullName,
		Email:      arg.Email,
		Policy:     arg.Policy,
		Active:     arg.Active,
		Malignancy: arg.Malignancy,
		BirthDate:  arg.BirthDate.Format(time.RFC3339),
	})
	if err != nil {
		return adapter_errors.HandleGRPCError(err)
	}
	return nil
}

func (a *adapter) GetPatient(ctx context.Context, id uuid.UUID) (domain.Patient, error) {
	res, err := a.client.GetPatient(ctx, &pb.GetPatientIn{
		Id: id.String(),
	})
	if err != nil {
		return domain.Patient{}, adapter_errors.HandleGRPCError(err)
	}

	return mappers.Patient{}.Domain(res.Patient), nil
}

func (a *adapter) GetPatientsByDoctorID(ctx context.Context, doctorID uuid.UUID) ([]domain.Patient, error) {
	res, err := a.client.GetPatientsByDoctorID(ctx, &pb.GetPatientsByDoctorIDIn{
		Id: doctorID.String(),
	})
	slog.Info("GetPatientsByDoctorID", "res", res, "err", err)
	if err != nil {
		return nil, adapter_errors.HandleGRPCError(err)
	}

	return mappers.Patient{}.SliceDomain(res.Patients), nil
}

func (a *adapter) UpdatePatient(ctx context.Context, arg UpdatePatientIn) (domain.Patient, error) {
	var lastUziDate *string
	if arg.LastUziDate != nil {
		date := arg.LastUziDate.Format(time.RFC3339)
		lastUziDate = &date
	}

	res, err := a.client.UpdatePatient(ctx, &pb.UpdatePatientIn{
		Id:          arg.Id.String(),
		Active:      arg.Active,
		Malignancy:  arg.Malignancy,
		LastUziDate: lastUziDate,
	})
	if err != nil {
		return domain.Patient{}, adapter_errors.HandleGRPCError(err)
	}

	return mappers.Patient{}.Domain(res.Patient), nil
}
