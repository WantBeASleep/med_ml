package doctor

import (
	"context"

	"med/internal/generated/grpc/service"
	"med/internal/services/doctor"

	"github.com/golang/protobuf/ptypes/empty"
)

type DoctorHandler interface {
	RegisterDoctor(ctx context.Context, in *service.RegisterDoctorIn) (*empty.Empty, error)
	GetDoctor(ctx context.Context, in *service.GetDoctorIn) (*service.GetDoctorOut, error)
}

type handler struct {
	doctorSrv doctor.Service
}

func New(
	doctorSrv doctor.Service,
) DoctorHandler {
	return &handler{
		doctorSrv: doctorSrv,
	}
}
