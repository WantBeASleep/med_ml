package chat

import (
	"context"

	"github.com/AlekSi/pointer"

	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/chat/mappers"
)

func (h *handler) ChatsChatidHistoryGet(ctx context.Context, params api.ChatsChatidHistoryGetParams) (api.ChatsChatidHistoryGetRes, error) {
	limit := 50
	if params.Limit.IsSet() {
		limit = params.Limit.Value
	}

	offset := 0
	if params.Offset.IsSet() {
		offset = params.Offset.Value
	}

	messages, err := h.services.ChatService.GetChatHistory(ctx, params.Chatid, limit, offset)
	if err != nil {
		return nil, err
	}

	return pointer.To(
		api.ChatsChatidHistoryGetOKApplicationJSON(
			mappers.Message{}.SliceApi(messages),
		),
	), nil
}
