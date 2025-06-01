package grpc

import "errors"

var (
	ErrChatNameRequired      = errors.New("chat name is required")
	ErrPatientIDRequired     = errors.New("patient_id is required")
	ErrChatIDRequired        = errors.New("chat_id is required")
	ErrDoctorIDRequired      = errors.New("doctor_id is required")
	ErrParticipantIDRequired = errors.New("participant_ids is required and must be at least 2")
)
