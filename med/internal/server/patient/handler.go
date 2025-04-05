package patient

import (
	"context"

	pb "med/internal/generated/grpc/service"
	"med/internal/services/patient"

	"github.com/golang/protobuf/ptypes/empty"
)

type PatientHandler interface {
	CreatePatient(ctx context.Context, in *pb.CreatePatientIn) (*empty.Empty, error)
	GetPatient(ctx context.Context, in *pb.GetPatientIn) (*pb.GetPatientOut, error)
	GetPatientsByDoctorID(ctx context.Context, in *pb.GetPatientsByDoctorIDIn) (*pb.GetPatientsByDoctorIDOut, error)
	UpdatePatient(ctx context.Context, in *pb.UpdatePatientIn) (*pb.UpdatePatientOut, error)
}

type handler struct {
	patientSrv patient.Service
}

func New(
	patientSrv patient.Service,
) PatientHandler {
	return &handler{
		patientSrv: patientSrv,
	}
}
