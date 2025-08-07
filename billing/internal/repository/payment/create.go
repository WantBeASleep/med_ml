package payment

import (
	"billing/internal/repository/payment/entity"
)

func (q *repo) InsertPayment(payment entity.Payment) error {
	query := q.QueryBuilder().
		Insert(paymentTable).
		Columns(
			columnID,
			columnUserID,
			columnSubscriptionID,
			columnAmount,
			columnStatus,
			columnPaymentProviderID,
			columnPspToken,
			columnCreatedAt,
			columnUpdatedAt,
		).
		Values(
			payment.ID,
			payment.UserID,
			payment.SubscriptionID,
			payment.Amount,
			payment.Status,
			payment.PaymentProviderID,
			payment.PspToken,
			payment.CreatedAt,
			payment.UpdatedAt,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return nil
}
