package mappers

import (
	"time"

	"github.com/google/uuid"

	domain "composition-api/internal/domain/chat"
	pb "composition-api/internal/generated/grpc/clients/chat"
)

type Chat struct{}

func (m Chat) Domain(pb *pb.Chat) domain.Chat {
	participantIDs := make([]uuid.UUID, len(pb.ParticipantIds))
	for i, idStr := range pb.ParticipantIds {
		participantIDs[i] = uuid.MustParse(idStr)
	}

	createdAt, _ := time.Parse(time.RFC3339, pb.CreatedAt)

	var description *string
	if pb.Description != "" {
		description = &pb.Description
	}

	return domain.Chat{
		ID:             uuid.MustParse(pb.Id),
		Name:           pb.Name,
		Description:    description,
		PatientID:      uuid.MustParse(pb.PatientId),
		ParticipantIDs: participantIDs,
		CreatedAt:      createdAt,
	}
}

func (m Chat) SliceDomain(pbs []*pb.Chat) []domain.Chat {
	chats := make([]domain.Chat, len(pbs))
	for i, pb := range pbs {
		chats[i] = m.Domain(pb)
	}
	return chats
}

type Message struct{}

func (m Message) Domain(pb *pb.Message) domain.Message {
	createdAt, _ := time.Parse(time.RFC3339, pb.CreatedAt)

	return domain.Message{
		ID:        uuid.MustParse(pb.Id),
		ChatID:    uuid.MustParse(pb.ChatId),
		SenderID:  uuid.MustParse(pb.SenderId),
		Content:   pb.Content,
		CreatedAt: createdAt,
	}
}

func (m Message) SliceDomain(pbs []*pb.Message) []domain.Message {
	messages := make([]domain.Message, len(pbs))
	for i, pb := range pbs {
		messages[i] = m.Domain(pb)
	}
	return messages
}
