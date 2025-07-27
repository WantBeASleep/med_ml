package entity

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MessageContentType string

const (
	MessageContentTypeChatMessage   MessageContentType = "chat_message"
	MessageContentTypeNotification  MessageContentType = "notification"
	MessageContentTypeSystemMessage MessageContentType = "system_message"
)

func (m MessageContentType) Validate() error {
	switch m {
	case MessageContentTypeChatMessage, MessageContentTypeNotification, MessageContentTypeSystemMessage:
		return nil
	}

	return fmt.Errorf("invalid content type: %s", m)
}

type MessageType string

const (
	MessageTypeIn  MessageType = "in"  // incoming message (to client)
	MessageTypeOut MessageType = "out" // outgoing message (from client)
)

func (m MessageType) Validate() error {
	switch m {
	case MessageTypeIn, MessageTypeOut:
		return nil
	}

	return fmt.Errorf("invalid message type: %s", m)
}

type Message struct {
	ID          uuid.UUID          `json:"id"`
	ChatID      uuid.UUID          `json:"chat_id"`
	SenderID    uuid.UUID          `json:"sender_id"`
	Content     string             `json:"content"`
	Type        MessageType        `json:"type"`
	ContentType MessageContentType `json:"content_type"`
	CreatedAt   time.Time          `json:"created_at"`
}

func (m *Message) Validate() error {
	if m.ChatID == uuid.Nil {
		return ErrEmptyChatID
	}

	if m.SenderID == uuid.Nil {
		return ErrEmptySenderID
	}

	if m.Content == "" {
		return ErrEmptyMessage
	}

	if err := m.ContentType.Validate(); err != nil {
		return fmt.Errorf("validate content type: %w", err)
	}

	if err := m.Type.Validate(); err != nil {
		return fmt.Errorf("validate message type: %w", err)
	}

	return nil
}

func (m *Message) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// TODO need correct unmarshalling!
func (m *Message) NewFromJSON(data []byte) error {
	if err := json.Unmarshal(data, m); err != nil {
		return fmt.Errorf("unmarshal message: %w", err)
	}

	m.ID = uuid.New()
	m.CreatedAt = time.Now()

	if err := m.Validate(); err != nil {
		return fmt.Errorf("validate message: %w", err)
	}

	return nil
}
