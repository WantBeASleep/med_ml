package chat

import (
	"context"

	api "composition-api/internal/generated/http/api"
	services "composition-api/internal/services"
)

type ChatHandler interface {
	ChatsChatidGet(ctx context.Context, params api.ChatsChatidGetParams) (api.ChatsChatidGetRes, error)
	ChatsGet(ctx context.Context, params api.ChatsGetParams) (api.ChatsGetRes, error)
	ChatsPost(ctx context.Context, req *api.ChatsPostReq) (api.ChatsPostRes, error)
	ChatsChatidHistoryGet(ctx context.Context, params api.ChatsChatidHistoryGetParams) (api.ChatsChatidHistoryGetRes, error)
}

type handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) ChatHandler {
	return &handler{
		services: services,
	}
}
