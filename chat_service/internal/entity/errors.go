package entity

import "errors"

var (
	ErrEmptyChatName           = errors.New("chat name is empty")
	ErrEmptyChatID             = errors.New("chat id is empty")
	ErrLessThanTwoParticipants = errors.New("participant ids must be at least 2")
	ErrEmptyParticipantID      = errors.New("participant id is empty")

	ErrEmptyDoctorID       = errors.New("doctor id is empty")
	ErrEmptyDoctorFullname = errors.New("doctor fullname is empty")

	ErrEmptyMessage   = errors.New("message cannot be empty")
	ErrEmptySenderID  = errors.New("sender ID cannot be empty")
	ErrEmptyCreatedAt = errors.New("created at cannot be empty")

	ErrEmptyPatientID       = errors.New("patient id is empty")
	ErrEmptyPatientFullname = errors.New("patient fullname is empty")
	ErrEmptyPatientEmail    = errors.New("patient email is empty")
	ErrEmptyPatientPolicy   = errors.New("patient policy is empty")

	ErrInvalidWSMessageType = errors.New("invalid websocket message type")
	ErrInvalidContentType   = errors.New("invalid content type")
	ErrEmptyWSContent       = errors.New("websocket content cannot be empty")
	ErrEmptyWSMessageID     = errors.New("websocket message id is empty")
)
