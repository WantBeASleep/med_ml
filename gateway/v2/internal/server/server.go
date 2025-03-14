package server

import (
	"context"
	"fmt"

	api "gateway/internal/generated/http/api"
	"gateway/internal/server/uzi"
	services "gateway/internal/services"
)

type server struct {
	uzi.UziRoute
}

func New(services *services.Services) api.Handler {
	uziRoute := uzi.NewUziRoute(services)

	return &server{
		UziRoute: uziRoute,
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
