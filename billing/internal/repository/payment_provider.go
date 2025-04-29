package repository

import (
	"billing/internal/repository/entity"

	"github.com/Masterminds/squirrel"
	"github.com/WantBeASleep/goooool/daolib"
	"github.com/google/uuid"
)

const paymentProviderTable = "payment_provider"

type PaymentProviderQuery interface {
	InsertPaymentProvider(provider entity.PaymentProvider) error
	GetPaymentProviderByID(id uuid.UUID) (entity.PaymentProvider, error)
	UpdatePaymentProvider(provider entity.PaymentProvider) (int64, error)
	GetPaymentProviderByName(name string) (entity.PaymentProvider, error)
	ListPaymentProviders() ([]entity.PaymentProvider, error)
}

type paymentProviderQuery struct {
	*daolib.BaseQuery
}

func (q *paymentProviderQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *paymentProviderQuery) InsertPaymentProvider(provider entity.PaymentProvider) error {
	query := q.QueryBuilder().
		Insert(paymentProviderTable).
		Columns(
			"id",
			"name",
			"is_active",
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

func (q *paymentProviderQuery) GetPaymentProviderByID(id uuid.UUID) (entity.PaymentProvider, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"name",
			"is_active",
		).
		From(paymentProviderTable).
		Where(squirrel.Eq{
			"id": id,
		})

	var provider entity.PaymentProvider
	if err := q.Runner().Getx(q.Context(), &provider, query); err != nil {
		return entity.PaymentProvider{}, err
	}

	return provider, nil
}

func (q *paymentProviderQuery) UpdatePaymentProvider(provider entity.PaymentProvider) (int64, error) {
	query := q.QueryBuilder().
		Update(paymentProviderTable).
		SetMap(squirrel.Eq{
			"name":      provider.Name,
			"is_active": provider.IsActive,
		}).
		Where(squirrel.Eq{
			"id": provider.ID,
		})

	res, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (q *paymentProviderQuery) GetPaymentProviderByName(name string) (entity.PaymentProvider, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"name",
			"is_active",
		).
		From(paymentProviderTable).
		Where(squirrel.Eq{
			"name": name,
		})

	var provider entity.PaymentProvider
	if err := q.Runner().Getx(q.Context(), &provider, query); err != nil {
		return entity.PaymentProvider{}, err
	}

	return provider, nil
}

func (q *paymentProviderQuery) ListPaymentProviders() ([]entity.PaymentProvider, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"name",
			"is_active",
		).
		From(paymentProviderTable)

	var providers []entity.PaymentProvider
	if err := q.Runner().Selectx(q.Context(), &providers, query); err != nil {
		return nil, err
	}

	return providers, nil
}
