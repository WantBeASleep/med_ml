package chat

import (
	"context"

	"github.com/AlekSi/pointer"

	api "composition-api/internal/generated/http/api"
)

func (h *handler) ChatsPost(ctx context.Context, req *api.ChatsPostReq) (api.ChatsPostRes, error) {
	var description string
	if req.Description.IsSet() {
		description = req.Description.Value
	}

	chatID, err := h.services.ChatService.CreateChat(ctx, req.Name, description, req.PatientID, req.ParticipantIds)
	if err != nil {
		return nil, err
	}

	return pointer.To(api.SimpleUuid{
		ID: chatID,
	}), nil
}
