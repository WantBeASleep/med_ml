package services

import (
	"composition-api/internal/adapters"
	"composition-api/internal/dbus/producers"
	"composition-api/internal/repository"
	"composition-api/internal/services/card"
	"composition-api/internal/services/cytology"
	"composition-api/internal/services/device"
	"composition-api/internal/services/doctor"
	"composition-api/internal/services/download"
	"composition-api/internal/services/image"
	"composition-api/internal/services/node"
	"composition-api/internal/services/node_segment"
	"composition-api/internal/services/patient"
	"composition-api/internal/services/payment_provider"
	"composition-api/internal/services/register"
	"composition-api/internal/services/segment"
	"composition-api/internal/services/subscription"
	"composition-api/internal/services/tariff_plan"
	"composition-api/internal/services/tiler"
	"composition-api/internal/services/tokens"
	"composition-api/internal/services/uzi"
	"composition-api/internal/services/yookassa_webhook"
)

type Services struct {
	DeviceService          device.Service
	UziService             uzi.Service
	ImageService           image.Service
	NodeService            node.Service
	SegmentService         segment.Service
	NodeSegmentService     node_segment.Service
	TokensService          tokens.Service
	CardService            card.Service
	DoctorService          doctor.Service
	PatientService         patient.Service
	RegisterService        register.Service
	DownloadService        download.Service
	TariffPlanService      tariff_plan.Service
	SubscriptionService    subscription.Service
	PaymentProviderService payment_provider.Service
	YookassaWebhookService yookassa_webhook.Service
	CytologyService        cytology.Service
	TilerService           tiler.Service
}

func New(
	adapters *adapters.Adapters,
	producers producers.Producer,
	dao repository.DAO,
) *Services {
	deviceService := device.New(adapters)
	uziService := uzi.New(adapters, dao, producers)
	imageService := image.New(adapters)
	nodeService := node.New(adapters)
	segmentService := segment.New(adapters)
	nodeSegmentService := node_segment.New(adapters)
	tokenService := tokens.New(adapters)
	cardService := card.New(adapters)
	doctorService := doctor.New(adapters)
	patientService := patient.New(adapters)
	registerService := register.New(adapters)
	cytologyService := cytology.New(adapters, dao, producers)
	downloadService := download.New(dao, cytologyService)
	tariffPlanService := tariff_plan.New(adapters)
	subscriptionService := subscription.New(adapters)
	paymentProviderService := payment_provider.New(adapters)
	yookassaWebhookService := yookassa_webhook.New(adapters)
	tilerService := tiler.New(adapters.Tiler)

	return &Services{
		DeviceService:          deviceService,
		UziService:             uziService,
		ImageService:           imageService,
		NodeService:            nodeService,
		SegmentService:         segmentService,
		NodeSegmentService:     nodeSegmentService,
		TokensService:          tokenService,
		CardService:            cardService,
		DoctorService:          doctorService,
		PatientService:         patientService,
		RegisterService:        registerService,
		DownloadService:        downloadService,
		TariffPlanService:      tariffPlanService,
		SubscriptionService:    subscriptionService,
		PaymentProviderService: paymentProviderService,
		YookassaWebhookService: yookassaWebhookService,
		CytologyService:        cytologyService,
		TilerService:           tilerService,
	}
}
