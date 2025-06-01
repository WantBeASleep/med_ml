package httpserver

import (
	"net/http"

	"chat_service/internal/controller/websocket"

	"github.com/google/uuid"
	"github.com/lxzan/gws"
	"github.com/rs/zerolog/log"
)

func WSToHandle(upgrader *gws.Upgrader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientID := uuid.New() // TODO: temporary

		socket, err := upgrader.Upgrade(w, r)
		if err != nil {
			log.Error().
				Err(err).
				Msg("upgrade connection")

			return
		}

		socket.Session().Store(websocket.ClientIDKey, clientID)

		log.Info().
			Str("client_id", clientID.String()).
			Msg("WS connection established")

		go func() {
			socket.ReadLoop()
		}()
	}
}
