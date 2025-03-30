package flow

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"

	"uzi/internal/domain"
	pb "uzi/internal/generated/grpc/service"
)

var UziInit flowfuncDepsInjector = func(deps *Deps) flowfunc {
	return func(ctx context.Context, data FlowData) (FlowData, error) {
		projection := gofakeit.Word()
		externalId := uuid.New()
		author := uuid.New()

		resp, err := deps.Adapter.CreateUzi(ctx, &pb.CreateUziIn{
			Projection: projection,
			ExternalId: externalId.String(),
			DeviceId:   int64(data.Device.Id),
			Author:     author.String(),
		})
		if err != nil {
			return FlowData{}, fmt.Errorf("create uzi: %w", err)
		}

		data.Uzi = domain.Uzi{
			Id:         uuid.MustParse(resp.Id),
			Projection: projection,
			ExternalID: externalId,
			Author:     author,
			DeviceID:   data.Device.Id,
		}
		return data, nil
	}
}
