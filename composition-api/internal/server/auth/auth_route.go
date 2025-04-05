package auth

import (
	tokens "composition-api/internal/server/auth/tokens"
	services "composition-api/internal/services"
)

type AuthRoute interface {
	tokens.TokensHandler
}

type authRoute struct {
	tokens.TokensHandler
}

func NewAuthRoute(services *services.Services) AuthRoute {
	tokensHandler := tokens.NewHandler(services)

	return &authRoute{
		TokensHandler: tokensHandler,
	}
}
