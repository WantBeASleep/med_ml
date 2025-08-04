package postgres

import (
	"time"

	"github.com/google/uuid"
)

type ChatDTO struct {
	ID           uuid.UUID `db:"id"`
	Name         string    `db:"name"`
	Description  string    `db:"description"`
	PatientID    uuid.UUID `db:"patient_id"`
	LastActivity time.Time `db:"last_activity"`
	CreatedAt    time.Time `db:"created_at"`
}

type MessageDTO struct {
	ID          uuid.UUID `db:"id"`
	ChatID      uuid.UUID `db:"chat_id"`
	SenderID    uuid.UUID `db:"sender_id"`
	Content     string    `db:"content"`
	Type        string    `db:"type"`
	ContentType string    `db:"content_type"`
	CreatedAt   time.Time `db:"created_at"`
}

type ChatParticipantDTO struct {
	ChatID   uuid.UUID `db:"chat_id"`
	DoctorID uuid.UUID `db:"doctor_id"`
}
