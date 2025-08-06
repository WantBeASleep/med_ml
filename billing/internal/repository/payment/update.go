package payment

import (
	"billing/internal/repository/payment/entity"

	sq "github.com/Masterminds/squirrel"
)

func (q *repo) UpdatePayment(payment entity.Payment) (int64, error) {
	query := q.QueryBuilder().
		Update(paymentTable).
		SetMap(sq.Eq{
			columnAmount:            payment.Amount,
			columnStatus:            payment.Status,
			columnPaymentProviderID: payment.PaymentProviderID,
			columnPspToken:          payment.PspToken,
			columnUpdatedAt:         payment.UpdatedAt,
		}).
		Where(sq.Eq{
			columnID: payment.ID,
		})

	res, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
