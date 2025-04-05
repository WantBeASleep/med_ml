package doctor

import (
	"context"

	"med/internal/domain"
	"med/internal/generated/grpc/service"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) RegisterDoctor(ctx context.Context, in *service.RegisterDoctorIn) (*empty.Empty, error) {
	doctorID, err := uuid.Parse(in.Doctor.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат ID врача: %s", err.Error())
	}

	if err := h.doctorSrv.RegisterDoctor(ctx, domain.Doctor{
		Id:          doctorID,
		FullName:    in.Doctor.Fullname,
		Org:         in.Doctor.Org,
		Job:         in.Doctor.Job,
		Description: in.Doctor.Description,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &empty.Empty{}, nil
}
