package payment_provider

import (
	"billing/internal/repository/payment_provider/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (q *repo) GetPaymentProviderByID(id uuid.UUID) (entity.PaymentProvider, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnName,
			columnIsActive,
		).
		From(table).
		Where(sq.Eq{
			columnID: id,
		})

	var provider entity.PaymentProvider
	if err := q.Runner().Getx(q.Context(), &provider, query); err != nil {
		return entity.PaymentProvider{}, err
	}

	return provider, nil
}

func (q *repo) GetPaymentProviderByName(name string) (entity.PaymentProvider, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnName,
			columnIsActive,
		).
		From(table).
		Where(sq.Eq{
			columnName: name,
		})

	var provider entity.PaymentProvider
	if err := q.Runner().Getx(q.Context(), &provider, query); err != nil {
		return entity.PaymentProvider{}, err
	}

	return provider, nil
}

func (q *repo) ListPaymentProviders() ([]entity.PaymentProvider, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnName,
			columnIsActive,
		).
		From(table)

	var providers []entity.PaymentProvider
	if err := q.Runner().Selectx(q.Context(), &providers, query); err != nil {
		return nil, err
	}

	return providers, nil
}
