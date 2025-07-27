package chat

import (
	"context"

	"github.com/google/uuid"

	"composition-api/internal/adapters/chat/mappers"
	domain "composition-api/internal/domain/chat"
	pb "composition-api/internal/generated/grpc/clients/chat"
)

type Adapter interface {
	CreateChat(ctx context.Context, name, description string, patientID uuid.UUID, participantIDs []uuid.UUID) (uuid.UUID, error)
	GetChats(ctx context.Context, doctorID uuid.UUID) ([]domain.Chat, error)
	GetChat(ctx context.Context, chatID uuid.UUID) (domain.Chat, error)
	GetChatHistory(ctx context.Context, chatID uuid.UUID, limit, offset int) ([]domain.Message, error)
}

type adapter struct {
	client pb.ChatSrvClient
}

func NewAdapter(client pb.ChatSrvClient) Adapter {
	return &adapter{
		client: client,
	}
}

func (a *adapter) CreateChat(ctx context.Context, name, description string, patientID uuid.UUID, participantIDs []uuid.UUID) (uuid.UUID, error) {
	participantIDStrs := make([]string, len(participantIDs))
	for i, id := range participantIDs {
		participantIDStrs[i] = id.String()
	}

	req := &pb.CreateChatIn{
		Name:           name,
		Description:    description,
		PatientId:      patientID.String(),
		ParticipantIds: participantIDStrs,
	}

	resp, err := a.client.CreateChat(ctx, req)
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.MustParse(resp.ChatId), nil
}

func (a *adapter) GetChats(ctx context.Context, doctorID uuid.UUID) ([]domain.Chat, error) {
	req := &pb.GetChatsIn{
		DoctorId: doctorID.String(),
	}

	resp, err := a.client.GetChats(ctx, req)
	if err != nil {
		return nil, err
	}

	return mappers.Chat{}.SliceDomain(resp.Chats), nil
}

func (a *adapter) GetChat(ctx context.Context, chatID uuid.UUID) (domain.Chat, error) {
	req := &pb.GetChatIn{
		ChatId: chatID.String(),
	}

	resp, err := a.client.GetChat(ctx, req)
	if err != nil {
		return domain.Chat{}, err
	}

	return mappers.Chat{}.Domain(resp.Chat), nil
}

func (a *adapter) GetChatHistory(ctx context.Context, chatID uuid.UUID, limit, offset int) ([]domain.Message, error) {
	req := &pb.GetChatHistoryIn{
		ChatId: chatID.String(),
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	resp, err := a.client.GetChatHistory(ctx, req)
	if err != nil {
		return nil, err
	}

	return mappers.Message{}.SliceDomain(resp.Messages), nil
}
