package med

import (
	"context"

	adapter_errors "composition-api/internal/adapters/errors"
	"composition-api/internal/adapters/med/mappers"
	domain "composition-api/internal/domain/med"
	pb "composition-api/internal/generated/grpc/clients/med"

	"github.com/google/uuid"
)

func (a *adapter) RegisterDoctor(ctx context.Context, doctor domain.Doctor) error {
	_, err := a.client.RegisterDoctor(ctx, &pb.RegisterDoctorIn{
		Doctor: &pb.Doctor{
			Id:          doctor.Id.String(),
			Fullname:    doctor.FullName,
			Org:         doctor.Org,
			Job:         doctor.Job,
			Description: doctor.Description,
		},
	})
	return err
}

func (a *adapter) GetDoctor(ctx context.Context, id uuid.UUID) (domain.Doctor, error) {
	res, err := a.client.GetDoctor(ctx, &pb.GetDoctorIn{
		Id: id.String(),
	})
	if err != nil {
		return domain.Doctor{}, adapter_errors.HandleGRPCError(err)
	}

	return mappers.Doctor{}.Domain(res.Doctor), nil
}
