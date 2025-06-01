package http_server

import (
	"chat_service/internal/config"
	"chat_service/internal/controller/websocket"

	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/lxzan/gws"
)

const (
	readBufferSize   = 32 * 1024
	writeBufferSize  = 32 * 1024
	maxPayloadSize   = 65535
	handshakeTimeout = 30 * time.Second
	parallelEnabled  = true
	parallelGolimit  = 1

	readTimeout  = 20 * time.Second
	writeTimeout = 10 * time.Second
	idleTimeout  = 20 * time.Second
)

func NewServer(cfg config.Config, chatService websocket.ChatService, messageService websocket.MessageService) *http.Server {
	upgrader := gws.NewUpgrader(websocket.NewWSHandler(cfg.WSClient, websocket.NewHub(chatService, messageService)), &gws.ServerOption{
		Recovery:            gws.Recovery,
		CheckUtf8Enabled:    false,
		ParallelEnabled:     parallelEnabled,
		ParallelGolimit:     parallelGolimit,
		WriteBufferSize:     writeBufferSize,
		ReadBufferSize:      readBufferSize,
		ReadMaxPayloadSize:  maxPayloadSize,
		WriteMaxPayloadSize: maxPayloadSize,
		HandshakeTimeout:    handshakeTimeout,
		TlsConfig: &tls.Config{
			VerifyConnection: func(state tls.ConnectionState) error { //nolint
				return nil
			},
			InsecureSkipVerify: true, //nolint
		},
		PermessageDeflate: gws.PermessageDeflate{
			Enabled: false,
		},
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", WSToHandle(upgrader))
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	//nolint:gosec
	return &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port),
		Handler:      mux,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}
}
