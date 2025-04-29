package server

import (
	"billing/internal/generated/grpc/service"
	"billing/internal/server/payment_provider"
	"billing/internal/server/subscription"
	"billing/internal/server/tariff_plan"
	"billing/internal/server/yookassa_webhook"
	"billing/internal/services"
)

type Handler struct {
	subscription.SubscriptionHandler
	tariff_plan.TariffPlanHandler
	yookassa_webhook.YookassaWebhookHandler
	payment_provider.PaymentProviderHandler

	service.UnsafeBillingServiceServer
}

func New(
	services *services.Services,
) *Handler {
	subscriptionHandler := subscription.New(services)
	tariffPlanHandler := tariff_plan.New(services)
	yookassaHandler := yookassa_webhook.New(services)
	paymentProviderHandler := payment_provider.New(services)

	return &Handler{
		SubscriptionHandler:    subscriptionHandler,
		TariffPlanHandler:      tariffPlanHandler,
		YookassaWebhookHandler: yookassaHandler,
		PaymentProviderHandler: paymentProviderHandler,
	}
}
