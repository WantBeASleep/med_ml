package chat_service

import (
	"context"
	"fmt"

	"chat_service/internal/entity"
	"chat_service/internal/repository/postgres"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type ChatRepository interface {
	CreateChatWithParticipants(ctx context.Context, chat postgres.ChatDTO, participantIDs []uuid.UUID) error
	GetChat(ctx context.Context, chatID uuid.UUID) (postgres.ChatDTO, error)
	GetChatParticipants(ctx context.Context, chatID uuid.UUID) ([]uuid.UUID, error)
	ListChatsByDoctorID(ctx context.Context, doctorID uuid.UUID) ([]postgres.ChatDTO, error)
	UpdateChatLastActivity(ctx context.Context, chatID uuid.UUID) error
}

type ChatService struct {
	chatRepo ChatRepository
}

func NewChatService(chatRepo ChatRepository) *ChatService {
	return &ChatService{
		chatRepo: chatRepo,
	}
}

func (s *ChatService) CreateChat(ctx context.Context, name, description string, patientID uuid.UUID, participantIDs []uuid.UUID) (uuid.UUID, error) {
	chat := entity.NewChat(name, description, patientID, participantIDs)

	chatDTO := postgres.ChatDTO{
		ID:           chat.ID,
		Name:         chat.Name,
		Description:  chat.Description,
		PatientID:    chat.PatientID,
		LastActivity: chat.LastActivity,
		CreatedAt:    chat.CreatedAt,
	}

	if err := s.chatRepo.CreateChatWithParticipants(ctx, chatDTO, participantIDs); err != nil {
		return uuid.UUID{}, fmt.Errorf("save chat to repository: %w", err)
	}

	log.Debug().
		Str("chatID", chat.ID.String()).
		Str("name", chat.Name).
		Int("participantsCount", len(participantIDs)).
		Msg("created new chat")

	return chat.ID, nil
}

func (s *ChatService) GetChat(ctx context.Context, chatID uuid.UUID) (entity.Chat, error) {
	chatDTO, err := s.chatRepo.GetChat(ctx, chatID)
	if err != nil {
		return entity.Chat{}, fmt.Errorf("get chat from repository: %w", err)
	}

	participantIDs, err := s.chatRepo.GetChatParticipants(ctx, chatID)
	if err != nil {
		return entity.Chat{}, fmt.Errorf("get chat participants: %w", err)
	}

	return entity.Chat{
		ID:             chatDTO.ID,
		Name:           chatDTO.Name,
		Description:    chatDTO.Description,
		PatientID:      chatDTO.PatientID,
		ParticipantIDs: participantIDs,
		LastActivity:   chatDTO.LastActivity,
		CreatedAt:      chatDTO.CreatedAt,
	}, nil
}

func (s *ChatService) ListChatsByDoctor(ctx context.Context, doctorID uuid.UUID) ([]entity.Chat, error) {
	chatsDTO, err := s.chatRepo.ListChatsByDoctorID(ctx, doctorID)
	if err != nil {
		return nil, fmt.Errorf("list chats from repository: %w", err)
	}

	var chats []entity.Chat

	for _, chatDTO := range chatsDTO {
		participantIDs, err := s.chatRepo.GetChatParticipants(ctx, chatDTO.ID)
		if err != nil {
			return nil, fmt.Errorf("get participants for chat %s: %w", chatDTO.ID, err)
		}

		chats = append(chats, entity.Chat{
			ID:             chatDTO.ID,
			Name:           chatDTO.Name,
			Description:    chatDTO.Description,
			PatientID:      chatDTO.PatientID,
			ParticipantIDs: participantIDs,
			LastActivity:   chatDTO.LastActivity,
			CreatedAt:      chatDTO.CreatedAt,
		})
	}

	return chats, nil
}

func (s *ChatService) UpdateChatLastActivity(ctx context.Context, chatID uuid.UUID) error {
	if err := s.chatRepo.UpdateChatLastActivity(ctx, chatID); err != nil {
		return fmt.Errorf("update chat last activity: %w", err)
	}

	return nil
}

func (s *ChatService) GetChatParticipants(ctx context.Context, chatID uuid.UUID) ([]uuid.UUID, error) {
	return s.chatRepo.GetChatParticipants(ctx, chatID)
}
