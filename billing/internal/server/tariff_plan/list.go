package tariff_plan

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "billing/internal/generated/grpc/service"

	"google.golang.org/protobuf/types/known/durationpb"
)

func (h *handler) ListTariffPlans(ctx context.Context, _ *empty.Empty) (*pb.ListTariffPlansOut, error) {
	tariffPlans, err := h.services.TariffPlan.ListTariffPlans(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list tariff plans: %s", err.Error())
	}

	var pbTariffPlans []*pb.TariffPlan
	for _, tp := range tariffPlans {
		pbTariffPlans = append(pbTariffPlans, &pb.TariffPlan{
			TariffPlanId: tp.ID.String(),
			Name:         tp.Name,
			Description:  tp.Description,
			Price:        tp.Price.String(),
			Duration:     durationpb.New(tp.Duration),
		})
	}

	return &pb.ListTariffPlansOut{TariffPlans: pbTariffPlans}, nil
}
