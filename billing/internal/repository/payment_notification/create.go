package payment_notification

import (
	"billing/internal/repository/payment_notification/entity"
)

func (q *repo) InsertPaymentNotification(notification entity.PaymentNotification) error {
	query := q.QueryBuilder().
		Insert(table).
		Columns(
			columnID,
			columnProviderPaymentID,
			columnEvent,
			columnPaymentProviderID,
			columnReceivedAt,
			columnNotificationData,
			columnIsValid,
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
