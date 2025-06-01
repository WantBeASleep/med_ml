package chat

import (
	"composition-api/internal/server/chat/chat"
	services "composition-api/internal/services"
)

type ChatRoute interface {
	chat.ChatHandler
}

type chatRoute struct {
	chat.ChatHandler
}

func NewChatRoute(services *services.Services) ChatRoute {
	chatHandler := chat.NewHandler(services)

	return &chatRoute{
		ChatHandler: chatHandler,
	}
}
