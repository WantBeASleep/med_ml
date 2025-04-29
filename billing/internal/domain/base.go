package domain

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

type SubscriptionStatus string

const (
	SubPendingPayment SubscriptionStatus = "pending_payment"
	SubActive         SubscriptionStatus = "active"
	SubCancelled      SubscriptionStatus = "cancelled"
)

type PaymentStatus string

const (
	PayPending           PaymentStatus = "pending"
	PayWaitingForCapture PaymentStatus = "waiting_for_capture"
	PayWaitingForCancel  PaymentStatus = "waiting_for_cancel"
	PayCompleted         PaymentStatus = "completed"
	PayCancelled         PaymentStatus = "pay_cancelled"
)

type PaymentProvider struct {
	ID       uuid.UUID
	Name     string
	IsActive bool
}

type TariffPlan struct {
	ID          uuid.UUID
	Name        string
	Description string
	Price       decimal.Decimal
	Duration    time.Duration
}

type Subscription struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	TariffPlanID uuid.UUID
	StartDate    time.Time
	EndDate      time.Time
	Status       SubscriptionStatus
}

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
