package flow

import (
	"context"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"

	"med/internal/domain"
	pb "med/internal/generated/grpc/service"
)

var CreatePatient flowfuncDepsInjector = func(deps *Deps) flowfunc {
	return func(ctx context.Context, data FlowData) (FlowData, error) {
		id := uuid.New()
		fullname := gofakeit.Name()
		email := gofakeit.Email()
		policy := gofakeit.MinecraftAnimal()
		active := gofakeit.Bool()
		malignancy := gofakeit.Bool()
		birthDate := gofakeit.Date().Round(0)

		_, err := deps.Adapter.CreatePatient(
			ctx,
			&pb.CreatePatientIn{
				Id:         id.String(),
				Fullname:   fullname,
				Email:      email,
				Policy:     policy,
				Active:     active,
				Malignancy: malignancy,
				BirthDate:  birthDate.Format(time.RFC3339),
			},
		)
		if err != nil {
			return FlowData{}, fmt.Errorf("create patient: %w", err)
		}

		data.Patient = domain.Patient{
			Id:         id,
			FullName:   fullname,
			Email:      email,
			Policy:     policy,
			Active:     active,
			Malignancy: malignancy,
			BirthDate:  birthDate,
		}
		return data, nil
	}
}
