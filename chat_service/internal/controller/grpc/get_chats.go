package grpc

import (
	"context"
	"fmt"

	chatv1 "chat_service/internal/pb/chat/v1"

	"github.com/google/uuid"
)

func (s *Server) GetChats(ctx context.Context, req *chatv1.GetChatsIn) (*chatv1.GetChatsOut, error) {
	if req.DoctorId == "" {
		return nil, ErrDoctorIDRequired
	}

	doctorID, err := uuid.Parse(req.DoctorId)
	if err != nil {
		return nil, fmt.Errorf("parse doctor ID: %w", err)
	}

	chats, err := s.chatUsecase.ListChats(ctx, doctorID)
	if err != nil {
		return nil, fmt.Errorf("list chats: %w", err)
	}

	var chatList []*chatv1.Chat

	for _, chat := range chats {
		var participantIDs []string

		for _, id := range chat.ParticipantIDs {
			participantIDs = append(participantIDs, id.String())
		}

		chatList = append(chatList, &chatv1.Chat{
			Id:             chat.ID.String(),
			Name:           chat.Name,
			Description:    chat.Description,
			PatientId:      chat.PatientID.String(),
			ParticipantIds: participantIDs,
			CreatedAt:      chat.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &chatv1.GetChatsOut{
		Chats: chatList,
	}, nil
}
