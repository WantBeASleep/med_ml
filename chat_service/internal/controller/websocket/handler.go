package websocket

import (
	"errors"
	"time"

	"chat_service/internal/config"
	"chat_service/internal/entity"

	"github.com/google/uuid"
	"github.com/lxzan/gws"
	"github.com/rs/zerolog/log"
)

const (
	ClientIDKey = "client_id"
)

type WSHandler interface {
	gws.Event
}

type wsHandler struct {
	pongWait     time.Duration
	pingInterval time.Duration

	hub *Hub
}

func (ws *wsHandler) OnOpen(socket *gws.Conn) {
	clientID, exists := getParamFromConn(socket, ClientIDKey)
	if !exists {
		if err := socket.NetConn().Close(); err != nil {
			log.Error().
				Err(err).
				Msg("close connection")
		}

		log.Error().
			Msg("missing clientID from connection")

		return
	}

	ws.hub.Register(socket, clientID)

	if err := errors.Join(
		socket.SetReadDeadline(time.Now().Add(ws.pongWait)),
		socket.SetWriteDeadline(time.Time{}), // reset
	); err != nil {
		log.Error().
			Err(err).
			Msg("set deadlines")
	}

	log.Info().
		Str("clientID", clientID).
		Msg("WS connection opened")

	go startPing(socket, ws.pingInterval)
}

func (ws *wsHandler) OnClose(socket *gws.Conn, _ error) {
	clientID, _ := getParamFromConn(socket, ClientIDKey)

	ws.hub.Unregister(clientID)

	log.Info().
		Str("clientID", clientID).
		Msg("WS connection closed")
}

func (ws *wsHandler) OnPing(socket *gws.Conn, payload []byte) {
	if err := socket.SetReadDeadline(time.Now().Add(ws.pongWait)); err != nil {
		log.Error().
			Err(err).
			Msg("set read deadline")

		return
	}

	if err := socket.WritePong(payload); err != nil {
		log.Error().
			Err(err).
			Msg("write ping")
	}
}

func (ws *wsHandler) OnPong(socket *gws.Conn, payload []byte) {
	clientID, exists := getParamFromConn(socket, ClientIDKey)
	if !exists {
		log.Error().
			Msg("missing clientID from connection")
	}

	pingTime, err := time.Parse(time.RFC3339, string(payload))
	if err != nil {
		log.Error().
			Err(err).
			Msg("parse ping time")
	}

	log.Debug().
		Str("ping-pong duration", time.Since(pingTime).String()).
		Str("clientID", clientID).
		Msg("on pong")

	if err := socket.SetReadDeadline(time.Now().Add(ws.pongWait)); err != nil {
		log.Error().
			Err(err).
			Msg("set read deadline")
	}
}

func (ws *wsHandler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()

	clientID, exists := getParamFromConn(socket, ClientIDKey)
	if !exists {
		log.Error().
			Msg("missing clientID from connection")

		return
	}

	var payload entity.Message
	if err := payload.NewFromJSON(message.Bytes()); err != nil {
		log.Error().
			Err(err).
			Msg("unmarshal message")

		return
	}

	log.Debug().
		Str("clientID", clientID).
		Msg("message received")

	clientIDUUID, err := uuid.Parse(clientID)
	if err != nil {
		log.Error().
			Err(err).
			Str("clientID", clientID).
			Msg("parse clientID")
	}

	payload.SenderID = clientIDUUID
	payload.ID = uuid.New()
	payload.Type = entity.MessageTypeOut
	payload.ContentType = entity.MessageContentTypeChatMessage
	payload.CreatedAt = time.Now()

	if err := payload.Validate(); err != nil {
		log.Error().
			Err(err).
			Str("clientID", clientID).
			Msg("validate message")

		return
	}

	if err := ws.hub.SaveMessage(payload); err != nil {
		log.Error().
			Err(err).
			Str("clientID", clientID).
			Msg("save message")

		return
	}

	msgJSON, err := payload.ToJSON()
	if err != nil {
		log.Error().
			Err(err).
			Str("clientID", clientID).
			Msg("marshal message")

		return
	}

	ws.hub.BroadcastToChat(payload.ChatID, msgJSON, socket)
}

func NewWSHandler(cfg config.WSClient, hub *Hub) WSHandler {
	return &wsHandler{
		pongWait:     cfg.PongWait,
		pingInterval: cfg.PingInterval,
		hub:          hub,
	}
}

func startPing(socket *gws.Conn, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		if err := socket.WritePing([]byte(time.Now().Format(time.RFC3339))); err != nil {
			return
		}
	}
}

func getParamFromConn(conn *gws.Conn, paramKey string) (string, bool) {
	paramAny, exists := conn.Session().Load(paramKey)
	if !exists {
		return "", false
	}

	paramValue, ok := paramAny.(string)
	if !ok {
		return "", false
	}

	return paramValue, true
}
