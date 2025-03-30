package flow

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"

	"med/internal/domain"
	pb "med/internal/generated/grpc/service"
)

var CreateDoctor flowfuncDepsInjector = func(deps *Deps) flowfunc {
	return func(ctx context.Context, data FlowData) (FlowData, error) {
		id := uuid.New()
		fullname := gofakeit.Name()
		org := gofakeit.Company()
		job := gofakeit.JobTitle()
		description := gofakeit.MinecraftAnimal()

		_, err := deps.Adapter.RegisterDoctor(
			ctx,
			&pb.RegisterDoctorIn{
				Doctor: &pb.Doctor{
					Id:          id.String(),
					Fullname:    fullname,
					Org:         org,
					Job:         job,
					Description: &description,
				},
			},
		)
		if err != nil {
			return FlowData{}, fmt.Errorf("create doctor: %w", err)
		}

		data.Doctor = domain.Doctor{
			Id:          id,
			FullName:    fullname,
			Org:         org,
			Job:         job,
			Description: &description,
		}
		return data, nil
	}
}
