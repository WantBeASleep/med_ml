package grpc

import (
	"context"
	"fmt"

	chatv1 "chat_service/internal/pb/chat/v1"

	"github.com/google/uuid"
)

func (s *Server) CreateChat(ctx context.Context, req *chatv1.CreateChatIn) (*chatv1.CreateChatOut, error) {
	if req.Name == "" {
		return nil, ErrChatNameRequired
	}

	if req.PatientId == "" {
		return nil, ErrPatientIDRequired
	}

	if len(req.ParticipantIds) < 2 {
		return nil, ErrParticipantIDRequired
	}

	patientID, err := uuid.Parse(req.PatientId)
	if err != nil {
		return nil, fmt.Errorf("parse patient ID: %w", err)
	}

	participantIDs := make([]uuid.UUID, len(req.ParticipantIds))

	for i, id := range req.ParticipantIds {
		participantID, err := uuid.Parse(id)
		if err != nil {
			return nil, fmt.Errorf("parse participant ID: %w", err)
		}

		participantIDs[i] = participantID
	}

	chatID, err := s.chatUsecase.CreateChat(
		ctx,
		req.Name,
		req.Description,
		patientID,
		participantIDs,
	)
	if err != nil {
		return nil, fmt.Errorf("create chat: %w", err)
	}

	return &chatv1.CreateChatOut{
		ChatId: chatID.String(),
	}, nil
}
