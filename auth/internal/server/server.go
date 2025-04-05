package server

import (
	"auth/internal/generated/grpc/service"
	"auth/internal/server/auth"
	"auth/internal/server/register"

	"auth/internal/services"
)

type Handler struct {
	auth.AuthHandler
	register.RegisterHandler

	service.UnsafeAuthSrvServer
}

func New(
	services *services.Services,
) *Handler {
	authHandler := auth.New(services)
	registerHandler := register.New(services)

	return &Handler{
		AuthHandler:     authHandler,
		RegisterHandler: registerHandler,
	}
}
