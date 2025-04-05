package flow

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v7"

	"med/internal/domain"
	pb "med/internal/generated/grpc/service"
)

var CreateCard flowfuncDepsInjector = func(deps *Deps) flowfunc {
	return func(ctx context.Context, data FlowData) (FlowData, error) {
		diagnosis := gofakeit.MinecraftAnimal()

		_, err := deps.Adapter.CreateCard(ctx, &pb.CreateCardIn{
			Card: &pb.Card{
				PatientId: data.Patient.Id.String(),
				DoctorId:  data.Doctor.Id.String(),
				Diagnosis: &diagnosis,
			},
		})
		if err != nil {
			return FlowData{}, fmt.Errorf("create card: %w", err)
		}

		data.Card = domain.Card{
			PatientID: data.Patient.Id,
			DoctorID:  data.Doctor.Id,
			Diagnosis: &diagnosis,
		}
		return data, nil
	}
}
