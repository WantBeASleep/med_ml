package repository

import (
	"context"

	"github.com/WantBeASleep/goooool/daolib"
	"github.com/jmoiron/sqlx"
)

type DAO interface {
	daolib.DAO
	NewTariffPlanQuery(ctx context.Context) TariffPlanQuery
	NewSubscriptionQuery(ctx context.Context) SubscriptionQuery
	NewPaymentQuery(ctx context.Context) PaymentQuery
	NewPaymentNotificationQuery(ctx context.Context) PaymentNotificationQuery
	NewPaymentProviderQuery(ctx context.Context) PaymentProviderQuery
}

type dao struct {
	daolib.DAO
}

func NewRepository(psql *sqlx.DB) DAO {
	return &dao{DAO: daolib.NewDao(psql)}
}

func (d *dao) NewTariffPlanQuery(ctx context.Context) TariffPlanQuery {
	tariffPlanQuery := &tariffPlanQuery{}
	d.NewRepo(ctx, tariffPlanQuery)

	return tariffPlanQuery
}

func (d *dao) NewSubscriptionQuery(ctx context.Context) SubscriptionQuery {
	subscriptionQuery := &subscriptionQuery{}
	d.NewRepo(ctx, subscriptionQuery)

	return subscriptionQuery
}

func (d *dao) NewPaymentQuery(ctx context.Context) PaymentQuery {
	paymentQuery := &paymentQuery{}
	d.NewRepo(ctx, paymentQuery)

	return paymentQuery
}

func (d *dao) NewPaymentNotificationQuery(ctx context.Context) PaymentNotificationQuery {
	paymentNotificationQuery := &paymentNotificationQuery{}
	d.NewRepo(ctx, paymentNotificationQuery)

	return paymentNotificationQuery
}

func (d *dao) NewPaymentProviderQuery(ctx context.Context) PaymentProviderQuery {
	paymentProviderQuery := &paymentProviderQuery{}
	d.NewRepo(ctx, paymentProviderQuery)

	return paymentProviderQuery
}
