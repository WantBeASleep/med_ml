package tokens

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type TokensHandler interface {
	LoginPost(ctx context.Context, req *api.LoginPostReq) (api.LoginPostRes, error)
	RefreshPost(ctx context.Context, req *api.RefreshPostReq) (api.RefreshPostRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) TokensHandler {
	return &handler{
		services: services,
	}
}
