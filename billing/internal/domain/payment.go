package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PaymentStatus string

const (
	PayPending           PaymentStatus = "pending"
	PayWaitingForCapture PaymentStatus = "waiting_for_capture"
	PayWaitingForCancel  PaymentStatus = "waiting_for_cancel"
	PayCompleted         PaymentStatus = "completed"
	PayCancelled         PaymentStatus = "pay_cancelled"
)

type Payment struct {
	ID                uuid.UUID
	UserID            uuid.UUID
	SubscriptionID    uuid.UUID
	Amount            decimal.Decimal
	Status            PaymentStatus
	PaymentProviderID uuid.UUID
	PspToken          string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type PaymentNotification struct {
	ID                uuid.UUID
	ProviderPaymentID string
	Event             string
	PaymentProviderID uuid.UUID
	ReceivedAt        time.Time
	NotificationData  map[string]interface{}
	IsValid           bool
}
