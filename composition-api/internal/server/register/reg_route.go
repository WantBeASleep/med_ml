package register

import (
	"composition-api/internal/server/register/register"
	services "composition-api/internal/services"
)

type RegisterRoute interface {
	register.RegisterHandler
}

type registerRoute struct {
	register.RegisterHandler
}

func NewRegisterRoute(services *services.Services) RegisterRoute {
	registerHandler := register.NewHandler(services)

	return &registerRoute{
		RegisterHandler: registerHandler,
	}
}
