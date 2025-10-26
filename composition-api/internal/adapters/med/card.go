package med

import (
	"context"
	"fmt"

	adapter_errors "composition-api/internal/adapters/errors"
	"composition-api/internal/adapters/med/mappers"
	domain "composition-api/internal/domain/med"
	pb "composition-api/internal/generated/grpc/clients/med"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *adapter) CreateCard(ctx context.Context, card domain.Card) error {
	_, err := a.client.CreateCard(ctx, &pb.CreateCardIn{
		Card: &pb.Card{
			DoctorId:  card.DoctorID.String(),
			PatientId: card.PatientID.String(),
			Diagnosis: card.Diagnosis,
		},
	})
	return err
}

func (a *adapter) GetCard(ctx context.Context, doctorID, patientID uuid.UUID) (domain.Card, error) {
	res, err := a.client.GetCard(ctx, &pb.GetCardIn{
		DoctorId:  doctorID.String(),
		PatientId: patientID.String(),
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return domain.Card{}, fmt.Errorf("unknown error: %w", err)
		}

		switch st.Code() {
		case codes.NotFound:
			return domain.Card{}, adapter_errors.ErrNotFound
		default:
			return domain.Card{}, err
		}
	}
	return mappers.Card{}.Domain(res.Card), nil
}

func (a *adapter) UpdateCard(ctx context.Context, card domain.Card) (domain.Card, error) {
	res, err := a.client.UpdateCard(ctx, &pb.UpdateCardIn{
		Card: &pb.Card{
			DoctorId:  card.DoctorID.String(),
			PatientId: card.PatientID.String(),
			Diagnosis: card.Diagnosis,
		},
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return domain.Card{}, fmt.Errorf("unknown error: %w", err)
		}

		switch st.Code() {
		case codes.NotFound:
			return domain.Card{}, adapter_errors.ErrNotFound
		default:
			return domain.Card{}, err
		}
	}

	return mappers.Card{}.Domain(res.Card), nil
}
