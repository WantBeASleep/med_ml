package grpc

import (
	"context"

	"chat_service/internal/entity"

	"github.com/google/uuid"
)

type ChatUsecase interface {
	CreateChat(ctx context.Context, name, description string, patientID uuid.UUID, participantIDs []uuid.UUID) (uuid.UUID, error)
	GetChat(ctx context.Context, chatID uuid.UUID) (entity.Chat, error)
	ListChats(ctx context.Context, doctorID uuid.UUID) ([]entity.Chat, error)
	GetChatHistory(ctx context.Context, chatID uuid.UUID, limit, offset int) ([]entity.Message, error)
}
