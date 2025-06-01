package grpc

import (
	"context"
	"fmt"

	chatv1 "chat_service/internal/pb/chat/v1"

	"github.com/google/uuid"
)

func (s *Server) GetChat(ctx context.Context, req *chatv1.GetChatIn) (*chatv1.GetChatOut, error) {
	if req.ChatId == "" {
		return nil, ErrChatIDRequired
	}

	chatID, err := uuid.Parse(req.ChatId)
	if err != nil {
		return nil, fmt.Errorf("parse chat ID: %w", err)
	}

	chat, err := s.chatUsecase.GetChat(ctx, chatID)
	if err != nil {
		return nil, fmt.Errorf("get chat: %w", err)
	}

	var participantIDs []string

	for _, id := range chat.ParticipantIDs {
		participantIDs = append(participantIDs, id.String())
	}

	return &chatv1.GetChatOut{
		Chat: &chatv1.Chat{
			Id:             chat.ID.String(),
			Name:           chat.Name,
			Description:    chat.Description,
			PatientId:      chat.PatientID.String(),
			ParticipantIds: participantIDs,
			CreatedAt:      chat.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	}, nil
}
