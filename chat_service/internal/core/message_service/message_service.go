package message_service

import (
	"context"
	"fmt"

	"chat_service/internal/entity"
	"chat_service/internal/repository/postgres"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type MessageRepository interface {
	GetMessages(ctx context.Context, chatID uuid.UUID, limit, offset int) ([]postgres.MessageDTO, error)
	SaveMessage(ctx context.Context, message postgres.MessageDTO) error
}

type CacheService interface {
	GetDoctor(ctx context.Context, doctorID string) (entity.Doctor, error)
	GetPatient(ctx context.Context, patientID string) (entity.Patient, error)
}

type MessageService struct {
	messageRepo  MessageRepository
	cacheService CacheService
}

func NewMessageService(messageRepo MessageRepository, cacheService CacheService) *MessageService {
	return &MessageService{
		messageRepo:  messageRepo,
		cacheService: cacheService,
	}
}

func (s *MessageService) GetMessages(ctx context.Context, chatID uuid.UUID, limit, offset int) ([]entity.Message, error) {
	messages, err := s.messageRepo.GetMessages(ctx, chatID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get messages from repository: %w", err)
	}

	log.Debug().
		Str("chatID", chatID.String()).
		Int("count", len(messages)).
		Int("limit", limit).
		Int("offset", offset).
		Msg("retrieved messages")

	var messagesEntity []entity.Message

	for _, message := range messages {
		messagesEntity = append(messagesEntity, entity.Message{
			ID:        message.ID,
			ChatID:    message.ChatID,
			SenderID:  message.SenderID,
			Content:   message.Content,
			CreatedAt: message.CreatedAt,
		})
	}

	return messagesEntity, nil
}

func (s *MessageService) SaveMessage(ctx context.Context, message entity.Message) error {
	if err := s.messageRepo.SaveMessage(ctx, postgres.MessageDTO{
		ID:        message.ID,
		ChatID:    message.ChatID,
		SenderID:  message.SenderID,
		Content:   message.Content,
		CreatedAt: message.CreatedAt,
	}); err != nil {
		return fmt.Errorf("save message: %w", err)
	}

	return nil
}
