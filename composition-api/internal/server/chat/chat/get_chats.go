package chat

import (
	"context"

	"github.com/AlekSi/pointer"

	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/chat/mappers"
)

func (h *handler) ChatsGet(ctx context.Context, params api.ChatsGetParams) (api.ChatsGetRes, error) {
	chats, err := h.services.ChatService.GetChats(ctx, params.DoctorID)
	if err != nil {
		return nil, err
	}

	return pointer.To(
		api.ChatsGetOKApplicationJSON(
			mappers.Chat{}.SliceApi(chats),
		),
	), nil
}
