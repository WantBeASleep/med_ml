package cytology_image

import (
	"context"
	"errors"

	"cytology/internal/domain"
	pb "cytology/internal/generated/grpc/service"
	"cytology/internal/server/mappers"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateCytologyImage(ctx context.Context, in *pb.CreateCytologyImageIn) (*pb.CreateCytologyImageOut, error) {
	externalID, err := uuid.Parse(in.ExternalId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "external_id is not a valid uuid: %s", err.Error())
	}

	doctorID, err := uuid.Parse(in.DoctorId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "doctor_id is not a valid uuid: %s", err.Error())
	}

	patientID, err := uuid.Parse(in.PatientId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "patient_id is not a valid uuid: %s", err.Error())
	}

	var prevID *uuid.UUID
	if in.PrevId != nil && *in.PrevId != "" {
		parsed, err := uuid.Parse(*in.PrevId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "prev_id is not a valid uuid: %s", err.Error())
		}
		prevID = &parsed
	}

	var parentPrevID *uuid.UUID
	if in.ParentPrevId != nil && *in.ParentPrevId != "" {
		parsed, err := uuid.Parse(*in.ParentPrevId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "parent_prev_id is not a valid uuid: %s", err.Error())
		}
		parentPrevID = &parsed
	}

	id, err := h.services.CytologyImage.CreateCytologyImage(ctx, mappers.CreateCytologyImageArgFromProto(in, externalID, doctorID, patientID, prevID, parentPrevID))
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrUnprocessableEntity):
			return nil, status.Errorf(codes.FailedPrecondition, "Ошибка валидации данных")
		default:
			return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
		}
	}

	return &pb.CreateCytologyImageOut{Id: id.String()}, nil
}
