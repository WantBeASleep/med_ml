package payment_provider

import (
	"billing/internal/repository/payment_provider/entity"

	sq "github.com/Masterminds/squirrel"
)

func (q *repo) UpdatePaymentProvider(provider entity.PaymentProvider) (int64, error) {
	query := q.QueryBuilder().
		Update(table).
		SetMap(sq.Eq{
			columnName:     provider.Name,
			columnIsActive: provider.IsActive,
		}).
		Where(sq.Eq{
			columnID: provider.ID,
		})

	res, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
