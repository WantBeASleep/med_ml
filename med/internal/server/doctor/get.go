package doctor

import (
	"context"
	"errors"

	"med/internal/generated/grpc/service"
	"med/internal/repository/entity"
	"med/internal/server/mappers"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) GetDoctor(ctx context.Context, in *service.GetDoctorIn) (*service.GetDoctorOut, error) {
	doctorID, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат ID врача: %s", err.Error())
	}

	doctor, err := h.doctorSrv.GetDoctor(ctx, doctorID)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			return nil, status.Errorf(codes.NotFound, "Врач не найден")
		default:
			return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
		}
	}

	return &service.GetDoctorOut{Doctor: mappers.DoctorFromDomain(doctor)}, nil
}
