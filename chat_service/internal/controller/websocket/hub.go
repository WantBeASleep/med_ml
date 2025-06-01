package websocket

import (
	"chat_service/internal/entity"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/lxzan/gws"
	"github.com/rs/zerolog/log"
)

const (
	timeout = 10 * time.Second
)

type ChatService interface {
	GetChatParticipants(ctx context.Context, chatID uuid.UUID) ([]uuid.UUID, error)
	UpdateChatLastActivity(ctx context.Context, chatID uuid.UUID) error
}

type MessageService interface {
	SaveMessage(ctx context.Context, message entity.Message) error
}

type Hub struct {
	mutex           sync.RWMutex
	userConnections map[string]*gws.Conn // user connections: key - userID, value - WebSocket connection

	chatService    ChatService
	messageService MessageService
}

func NewHub(chatService ChatService, messageService MessageService) *Hub {
	return &Hub{
		userConnections: make(map[string]*gws.Conn),
		chatService:     chatService,
		messageService:  messageService,
	}
}

func (h *Hub) Register(conn *gws.Conn, clientID string) {
	h.mutex.Lock()

	if oldConn, exists := h.userConnections[clientID]; exists {
		log.Info().
			Str("clientID", clientID).
			Msg("closing old connection for user")

		if err := oldConn.NetConn().Close(); err != nil {
			log.Error().
				Err(err).
				Str("clientID", clientID).
				Msg("close old connection")
		}
	}

	h.userConnections[clientID] = conn

	h.mutex.Unlock()

	log.Debug().
		Str("clientID", clientID).
		Int("totalConnections", len(h.userConnections)).
		Msg("client registered")
}

func (h *Hub) Unregister(clientID string) {
	h.mutex.Lock()
	delete(h.userConnections, clientID)
	defer h.mutex.Unlock()

	log.Debug().
		Str("clientID", clientID).
		Int("totalConnections", len(h.userConnections)).
		Msg("client unregistered")
}

func (h *Hub) SaveMessage(message entity.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := h.messageService.SaveMessage(ctx, message); err != nil {
		return fmt.Errorf("save message: %w", err)
	}

	return nil
}

func (h *Hub) BroadcastToChat(chatID uuid.UUID, message []byte, sender *gws.Conn) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	clients, err := h.chatService.GetChatParticipants(ctx, chatID)
	if err != nil {
		log.Error().
			Err(err).
			Str("chatID", chatID.String()).
			Msg("get chat participants")
		return
	}

	h.mutex.RLock()

	for _, clientID := range clients {
		conn, exists := h.userConnections[clientID.String()]
		if !exists {
			continue
		}

		if err := conn.WriteMessage(gws.OpcodeText, message); err != nil {
			log.Error().
				Err(err).
				Str("clientID", clientID.String()).
				Str("chatID", chatID.String()).
				Msg("send message to client")
			continue
		}

		// TODO: send notification to clients (when user is offline - save and onOpen will send it!)

		log.Debug().
			Str("clientID", clientID.String()).
			Str("chatID", chatID.String()).
			Msg("Message sent to user")
	}

	h.mutex.RUnlock()

	if err := h.chatService.UpdateChatLastActivity(ctx, chatID); err != nil {
		log.Error().
			Err(err).
			Str("chatID", chatID.String()).
			Msg("update chat last activity")
		return
	}
}
