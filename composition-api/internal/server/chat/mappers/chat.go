package mappers

import (
	domain "composition-api/internal/domain/chat"
	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/mappers"
)

type Chat struct{}

func (m Chat) Api(d domain.Chat) api.Chat {
	return api.Chat{
		ID:             d.ID,
		Name:           d.Name,
		Description:    mappers.ToOptString(d.Description),
		PatientID:      d.PatientID,
		ParticipantIds: d.ParticipantIDs,
		CreatedAt:      d.CreatedAt,
	}
}

func (m Chat) SliceApi(d []domain.Chat) []api.Chat {
	chats := make([]api.Chat, len(d))
	for i, chat := range d {
		chats[i] = m.Api(chat)
	}
	return chats
}

type Message struct{}

func (m Message) Api(d domain.Message) api.Message {
	return api.Message{
		ID:        d.ID,
		ChatID:    d.ChatID,
		SenderID:  d.SenderID,
		Content:   d.Content,
		CreatedAt: d.CreatedAt,
	}
}

func (m Message) SliceApi(d []domain.Message) []api.Message {
	messages := make([]api.Message, len(d))
	for i, message := range d {
		messages[i] = m.Api(message)
	}
	return messages
}
