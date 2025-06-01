package chat

import (
	"context"

	"github.com/AlekSi/pointer"

	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/chat/mappers"
)

func (h *handler) ChatsChatidGet(ctx context.Context, params api.ChatsChatidGetParams) (api.ChatsChatidGetRes, error) {
	chat, err := h.services.ChatService.GetChat(ctx, params.Chatid)
	if err != nil {
		return nil, err
	}

	return pointer.To(mappers.Chat{}.Api(chat)), nil
}
