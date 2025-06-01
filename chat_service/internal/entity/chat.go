package entity

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID             uuid.UUID
	Name           string
	Description    string
	PatientID      uuid.UUID
	ParticipantIDs []uuid.UUID
	LastActivity   time.Time
	CreatedAt      time.Time
}

func (c Chat) Validate() error {
	if c.ID == uuid.Nil {
		return ErrEmptyChatID
	}

	if c.Name == "" {
		return ErrEmptyChatName
	}

	if c.PatientID == uuid.Nil {
		return ErrEmptyPatientID
	}

	if len(c.ParticipantIDs) < 2 {
		return ErrLessThanTwoParticipants
	}

	for _, participantID := range c.ParticipantIDs {
		if participantID == uuid.Nil {
			return ErrEmptyParticipantID
		}
	}

	return nil
}

func NewChat(name, description string, patientID uuid.UUID, participantIDs []uuid.UUID) Chat {
	return Chat{
		ID:             uuid.New(),
		Name:           name,
		Description:    description,
		PatientID:      patientID,
		ParticipantIDs: participantIDs,
		CreatedAt:      time.Now(),
		LastActivity:   time.Now(),
	}
}
