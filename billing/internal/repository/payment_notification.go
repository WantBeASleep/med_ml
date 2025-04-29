package repository

import (
	"billing/internal/repository/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/WantBeASleep/goooool/daolib"
	"github.com/google/uuid"
)

const paymentNotificationTable = "payment_notification"

type PaymentNotificationQuery interface {
	InsertPaymentNotification(notification entity.PaymentNotification) error
	GetPaymentNotificationByID(id uuid.UUID) (entity.PaymentNotification, error)
}

type paymentNotificationQuery struct {
	*daolib.BaseQuery
}

func (q *paymentNotificationQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *paymentNotificationQuery) InsertPaymentNotification(notification entity.PaymentNotification) error {
	query := q.QueryBuilder().
		Insert(paymentNotificationTable).
		Columns(
			"id",
			"provider_payment_id",
			"event",
			"payment_provider_id",
			"received_at",
			"notification_data",
			"is_valid",
		).
		Values(
			notification.ID,
			notification.ProviderPaymentID,
			notification.Event,
			notification.PaymentProviderID,
			notification.ReceivedAt,
			notification.NotificationData,
			notification.IsValid,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return nil
}

func (q *paymentNotificationQuery) GetPaymentNotificationByID(id uuid.UUID) (entity.PaymentNotification, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"provider_payment_id",
			"event",
			"payment_provider_id",
			"received_at",
			"notification_data",
			"is_valid",
		).
		From(paymentNotificationTable).
		Where(sq.Eq{
			"id": id,
		})

	var notification entity.PaymentNotification
	if err := q.Runner().Getx(q.Context(), &notification, query); err != nil {
		return entity.PaymentNotification{}, err
	}

	return notification, nil
}
