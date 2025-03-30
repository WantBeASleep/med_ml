package flow

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"

	"uzi/internal/domain"
	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/server/mappers"
)

var UziInit flowfuncDepsInjector = func(deps *Deps) flowfunc {
	return func(ctx context.Context, data FlowData) (FlowData, error) {
		projection := domain.UziProjectionLong
		externalId := uuid.New()
		author := uuid.New()
		description := gofakeit.Word()

		resp, err := deps.Adapter.CreateUzi(ctx, &pb.CreateUziIn{
			Projection:  mappers.UziProjectionMap[projection],
			ExternalId:  externalId.String(),
			DeviceId:    int64(data.Device.Id),
			Author:      author.String(),
			Description: &description,
		})
		if err != nil {
			return FlowData{}, fmt.Errorf("create uzi: %w", err)
		}

		data.Uzi = domain.Uzi{
			Id:          uuid.MustParse(resp.Id),
			Projection:  projection,
			ExternalID:  externalId,
			Author:      author,
			DeviceID:    data.Device.Id,
			Description: &description,
		}
		return data, nil
	}
}
