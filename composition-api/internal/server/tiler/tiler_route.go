package tiler

import (
	tilerHandler "composition-api/internal/server/tiler/tiler"
	services "composition-api/internal/services"
)

type TilerRoute interface {
	tilerHandler.TilerHandler
}

type tilerRoute struct {
	tilerHandler.TilerHandler
}

func NewTilerRoute(services *services.Services) TilerRoute {
	handler := tilerHandler.NewHandler(services)

	return &tilerRoute{
		TilerHandler: handler,
	}
}
