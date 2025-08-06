package payment_notification

import (
	"billing/internal/repository/payment_notification/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (q *repo) GetPaymentNotificationByID(id uuid.UUID) (entity.PaymentNotification, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnProviderPaymentID,
			columnEvent,
			columnPaymentProviderID,
			columnReceivedAt,
			columnNotificationData,
			columnIsValid,
		).
		From(table).
		Where(sq.Eq{
			columnID: id,
		})

	var notification entity.PaymentNotification
	if err := q.Runner().Getx(q.Context(), &notification, query); err != nil {
		return entity.PaymentNotification{}, err
	}

	return notification, nil
}
