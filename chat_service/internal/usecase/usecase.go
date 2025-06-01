package usecase

import (
	"context"
	"fmt"

	"chat_service/internal/entity"

	"github.com/google/uuid"
)

type ChatService interface {
	CreateChat(ctx context.Context, name, description string, patientID uuid.UUID, participantIDs []uuid.UUID) (uuid.UUID, error)
	GetChat(ctx context.Context, chatID uuid.UUID) (entity.Chat, error)
	ListChatsByDoctor(ctx context.Context, doctorID uuid.UUID) ([]entity.Chat, error)
}

type MessageService interface {
	GetMessages(ctx context.Context, chatID uuid.UUID, limit, offset int) ([]entity.Message, error)
}

type ChatUsecase struct {
	chatService    ChatService
	messageService MessageService
}

func NewChatUsecase(
	chatService ChatService,
	messageService MessageService,
) *ChatUsecase {
	return &ChatUsecase{
		chatService:    chatService,
		messageService: messageService,
	}
}

func (u *ChatUsecase) CreateChat(ctx context.Context, name, description string, patientID uuid.UUID, participantIDs []uuid.UUID) (uuid.UUID, error) {
	chatID, err := u.chatService.CreateChat(ctx, name, description, patientID, participantIDs)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("create chat: %w", err)
	}

	return chatID, nil
}

func (u *ChatUsecase) GetChat(ctx context.Context, chatID uuid.UUID) (entity.Chat, error) {
	chat, err := u.chatService.GetChat(ctx, chatID)
	if err != nil {
		return entity.Chat{}, fmt.Errorf("get chat: %w", err)
	}

	return chat, nil
}

func (u *ChatUsecase) ListChats(ctx context.Context, doctorID uuid.UUID) ([]entity.Chat, error) {
	chats, err := u.chatService.ListChatsByDoctor(ctx, doctorID)
	if err != nil {
		return nil, fmt.Errorf("list chats: %w", err)
	}

	return chats, nil
}

func (u *ChatUsecase) GetChatHistory(ctx context.Context, chatID uuid.UUID, limit, offset int) ([]entity.Message, error) {
	messages, err := u.messageService.GetMessages(ctx, chatID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get chat history: %w", err)
	}

	return messages, nil
}
