package grpc

import (
	"context"
	"fmt"

	chatv1 "chat_service/internal/pb/chat/v1"

	"github.com/google/uuid"
)

func (s *Server) GetChatHistory(ctx context.Context, req *chatv1.GetChatHistoryIn) (*chatv1.GetChatHistoryOut, error) {
	if req.ChatId == "" {
		return nil, ErrChatIDRequired
	}

	chatID, err := uuid.Parse(req.ChatId)
	if err != nil {
		return nil, fmt.Errorf("parse chat ID: %w", err)
	}

	messages, err := s.chatUsecase.GetChatHistory(ctx, chatID, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, fmt.Errorf("get chat history: %w", err)
	}

	var messageList []*chatv1.Message

	for _, msg := range messages {
		messageList = append(messageList, &chatv1.Message{
			Id:        msg.ID.String(),
			ChatId:    msg.ChatID.String(),
			SenderId:  msg.SenderID.String(),
			Content:   msg.Content,
			CreatedAt: msg.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &chatv1.GetChatHistoryOut{
		Messages: messageList,
	}, nil
}
