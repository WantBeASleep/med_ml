package server

import (
	"context"
	"fmt"

	"composition-api/internal/server/billing"

	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/auth"
	"composition-api/internal/server/download"
	"composition-api/internal/server/med"
	"composition-api/internal/server/register"
	"composition-api/internal/server/uzi"
	services "composition-api/internal/services"
)

type server struct {
	auth.AuthRoute
	uzi.UziRoute
	med.MedRoute
	register.RegisterRoute
	download.DownloadRoute
	billing.BillingRoute
}

func New(services *services.Services) api.Handler {
	uziRoute := uzi.NewUziRoute(services)
	authRoute := auth.NewAuthRoute(services)
	medRoute := med.NewMedRoute(services)
	registerRoute := register.NewRegisterRoute(services)
	downloadRoute := download.NewDownloadRoute(services)
	billingRoute := billing.NewBillingRoute(services)

	return &server{
		UziRoute:      uziRoute,
		AuthRoute:     authRoute,
		MedRoute:      medRoute,
		RegisterRoute: registerRoute,
		DownloadRoute: downloadRoute,
		BillingRoute:  billingRoute,
	}
}

func (s *server) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	return &api.ErrorStatusCode{
		Response: api.Error{
			Message: fmt.Sprint("Необработанная ошибка сервера: ", err.Error()),
		},
	}
}
