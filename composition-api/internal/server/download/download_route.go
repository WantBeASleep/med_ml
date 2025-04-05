package download

import (
	download "composition-api/internal/server/download/download"
	services "composition-api/internal/services"
)

type DownloadRoute interface {
	download.DownloadHandler
}

type downloadRoute struct {
	download.DownloadHandler
}

func NewDownloadRoute(services *services.Services) DownloadRoute {
	handler := download.NewHandler(services)

	return &downloadRoute{
		DownloadHandler: handler,
	}
}
