package entity

import (
	"encoding/json"
	"time"

	"billing/internal/domain"

	"github.com/google/uuid"
)

type PaymentNotification struct {
	ID                uuid.UUID `db:"id"`
	ProviderPaymentID string    `db:"payment_id"`
	Event             string    `db:"event"`
	PaymentProviderID uuid.UUID `db:"payment_provider_id"`
	ReceivedAt        time.Time `db:"received_at"`
	NotificationData  string    `db:"notification_data"`
	IsValid           bool      `db:"is_valid"`
}

func (PaymentNotification) FromDomain(p domain.PaymentNotification) PaymentNotification {
	data, _ := json.Marshal(p.NotificationData)
	return PaymentNotification{
		ID:                p.ID,
		ProviderPaymentID: p.ProviderPaymentID,
		Event:             p.Event,
		PaymentProviderID: p.PaymentProviderID,
		ReceivedAt:        p.ReceivedAt,
		NotificationData:  string(data),
		IsValid:           p.IsValid,
	}
}

func (p PaymentNotification) ToDomain() domain.PaymentNotification {
	var data map[string]interface{}
	_ = json.Unmarshal([]byte(p.NotificationData), &data)
	return domain.PaymentNotification{
		ID:                p.ID,
		ProviderPaymentID: p.ProviderPaymentID,
		Event:             p.Event,
		PaymentProviderID: p.PaymentProviderID,
		ReceivedAt:        p.ReceivedAt,
		NotificationData:  data,
		IsValid:           p.IsValid,
	}
}
