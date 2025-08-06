package payment_provider

import (
	"billing/internal/repository/payment_provider/entity"
)

func (q *repo) InsertPaymentProvider(provider entity.PaymentProvider) error {
	query := q.QueryBuilder().
		Insert(table).
		Columns(
			columnID,
			columnName,
			columnIsActive,
		).
		Values(
			provider.ID,
			provider.Name,
			provider.IsActive,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return nil
}
