package tariff_plan

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "billing/internal/generated/grpc/service"

	"google.golang.org/protobuf/types/known/durationpb"
)

func (h *handler) GetTariffPlanByID(ctx context.Context, in *pb.GetTariffPlanByIDIn) (*pb.GetTariffPlanByIDOut, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid ID: %s", err.Error())
	}

	tariffPlan, err := h.services.TariffPlan.GetTariffPlanByID(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "tariff plan not found: %s", err.Error())
	}

	tariffPlanProto := &pb.TariffPlan{
		TariffPlanId: tariffPlan.ID.String(),
		Name:         tariffPlan.Name,
		Description:  tariffPlan.Description,
		Price:        tariffPlan.Price.String(),
		Duration:     durationpb.New(tariffPlan.Duration),
	}

	return &pb.GetTariffPlanByIDOut{
		TariffPlan: tariffPlanProto,
	}, nil
}
