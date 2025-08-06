package payment_notification

import (
	"billing/internal/repository/payment_notification/entity"

	daolib "github.com/WantBeASleep/goooool/daolib"
	"github.com/google/uuid"
)

const (
	table = "payment_notification"

	columnID                = "id"
	columnProviderPaymentID = "provider_payment_id"
	columnEvent             = "event"
	columnPaymentProviderID = "payment_provider_id"
	columnReceivedAt        = "received_at"
	columnNotificationData  = "notification_data"
	columnIsValid           = "is_valid"
)

type Repository interface {
	InsertPaymentNotification(notification entity.PaymentNotification) error
	GetPaymentNotificationByID(id uuid.UUID) (entity.PaymentNotification, error)
}

type repo struct {
	*daolib.BaseQuery
}

func NewRepo() *repo {
	return &repo{}
}

func (q *repo) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}
