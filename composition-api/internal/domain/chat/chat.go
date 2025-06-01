package chat

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID             uuid.UUID
	Name           string
	Description    *string
	PatientID      uuid.UUID
	ParticipantIDs []uuid.UUID
	CreatedAt      time.Time
}

type Message struct {
	ID        uuid.UUID
	ChatID    uuid.UUID
	SenderID  uuid.UUID
	Content   string
	CreatedAt time.Time
}
