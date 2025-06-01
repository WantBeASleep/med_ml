package server

import (
	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/auth"
	"composition-api/internal/server/chat"
	"composition-api/internal/server/download"
	"composition-api/internal/server/med"
	"composition-api/internal/server/register"
	"composition-api/internal/server/uzi"
	services "composition-api/internal/services"
	"context"
	"fmt"
)

type server struct {
	auth.AuthRoute
	uzi.UziRoute
	med.MedRoute
	register.RegisterRoute
	download.DownloadRoute
	chat.ChatRoute
}

func New(services *services.Services) api.Handler {
	uziRoute := uzi.NewUziRoute(services)
	authRoute := auth.NewAuthRoute(services)
	medRoute := med.NewMedRoute(services)
	registerRoute := register.NewRegisterRoute(services)
	downloadRoute := download.NewDownloadRoute(services)
	chatRoute := chat.NewChatRoute(services)

	return &server{
		UziRoute:      uziRoute,
		AuthRoute:     authRoute,
		MedRoute:      medRoute,
		RegisterRoute: registerRoute,
		DownloadRoute: downloadRoute,
		ChatRoute:     chatRoute,
	}
}

func (s *server) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	return &api.ErrorStatusCode{
		StatusCode: 500,
		Response: api.Error{
			Code:    500,
			Message: fmt.Sprint("Необработанная ошибка сервера: ", err.Error()),
		},
	}
}
