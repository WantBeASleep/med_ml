package chat

import (
	"context"

	"github.com/google/uuid"

	"composition-api/internal/adapters"
	domain "composition-api/internal/domain/chat"
)

type Service interface {
	CreateChat(ctx context.Context, name, description string, patientID uuid.UUID, participantIDs []uuid.UUID) (uuid.UUID, error)
	GetChats(ctx context.Context, doctorID uuid.UUID) ([]domain.Chat, error)
	GetChat(ctx context.Context, chatID uuid.UUID) (domain.Chat, error)
	GetChatHistory(ctx context.Context, chatID uuid.UUID, limit, offset int) ([]domain.Message, error)
}

type service struct {
	adapters *adapters.Adapters
}

func New(adapters *adapters.Adapters) Service {
	return &service{
		adapters: adapters,
	}
}

func (s *service) CreateChat(ctx context.Context, name, description string, patientID uuid.UUID, participantIDs []uuid.UUID) (uuid.UUID, error) {
	return s.adapters.Chat.CreateChat(ctx, name, description, patientID, participantIDs)
}

func (s *service) GetChats(ctx context.Context, doctorID uuid.UUID) ([]domain.Chat, error) {
	return s.adapters.Chat.GetChats(ctx, doctorID)
}

func (s *service) GetChat(ctx context.Context, chatID uuid.UUID) (domain.Chat, error) {
	return s.adapters.Chat.GetChat(ctx, chatID)
}

func (s *service) GetChatHistory(ctx context.Context, chatID uuid.UUID, limit, offset int) ([]domain.Message, error) {
	return s.adapters.Chat.GetChatHistory(ctx, chatID, limit, offset)
}
