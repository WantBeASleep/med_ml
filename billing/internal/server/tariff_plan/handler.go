package tariff_plan

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	pb "billing/internal/generated/grpc/service"
	"billing/internal/services"
)

type TariffPlanHandler interface {
	// CreateTariffPlan(ctx context.Context, req *pb.CreateTariffPlanIn) (*pb.CreateTariffPlanOut, error)
	GetTariffPlanByID(ctx context.Context, req *pb.GetTariffPlanByIDIn) (*pb.GetTariffPlanByIDOut, error)
	// UpdateTariffPlan(ctx context.Context, req *pb.UpdateTariffPlanIn) (*pb.UpdateTariffPlanOut, error)
	// DeleteTariffPlan(ctx context.Context, req *pb.DeleteTariffPlanIn) (*pb.DeleteTariffPlanOut, error)
	ListTariffPlans(ctx context.Context, _ *empty.Empty) (*pb.ListTariffPlansOut, error)
}
type handler struct {
	services *services.Services
}

func New(
	services *services.Services,
) TariffPlanHandler {
	return &handler{
		services: services,
	}
}
