package billing

import (
	"composition-api/internal/server/billing/payment_provider"
	"composition-api/internal/server/billing/subscription"
	"composition-api/internal/server/billing/tariff_plan"
	"composition-api/internal/server/billing/yookassa_webhook"
	services "composition-api/internal/services"
)

type BillingRoute interface {
	tariff_plan.TariffPlanHandler
	subscription.SubscriptionHandler
	payment_provider.PaymentProviderHandler
	yookassa_webhook.YookassaWebhookHandler
}

type billingRoute struct {
	tariff_plan.TariffPlanHandler
	subscription.SubscriptionHandler
	payment_provider.PaymentProviderHandler
	yookassa_webhook.YookassaWebhookHandler
}

func NewBillingRoute(services *services.Services) BillingRoute {
	tariffPlanHandler := tariff_plan.NewHandler(services)
	subscriptionHandler := subscription.NewHandler(services)
	paymentProviderHandler := payment_provider.NewHandler(services)
	yookassaWebhookHandler := yookassa_webhook.NewHandler(services)

	return &billingRoute{
		TariffPlanHandler:      tariffPlanHandler,
		SubscriptionHandler:    subscriptionHandler,
		PaymentProviderHandler: paymentProviderHandler,
		YookassaWebhookHandler: yookassaWebhookHandler,
	}
}
