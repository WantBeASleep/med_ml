package services

import (
	"log/slog"

	"auth/internal/services/auth"
	"auth/internal/services/password"
	"auth/internal/services/token"
	"auth/internal/services/user"

	"auth/internal/config"

	"auth/internal/repository"
)

type Services struct {
	AuthService     auth.Service
	PasswordService password.Service
	UserService     user.Service
	TokenService    token.Service
}

func New(
	dao repository.DAO,
	cfg *config.Config,
) *Services {
	publicKey, privateKey, err := cfg.ParseRsaKeys()
	if err != nil {
		slog.Error("parse rsa keys", "err", err)
		panic(err)
	}

	passwordService := password.New()
	tokenService := token.New(
		cfg.JWT.AccessTokenTime,
		cfg.JWT.RefreshTokenTime,
		privateKey,
		publicKey,
	)
	authService := auth.New(dao, passwordService, tokenService)
	userService := user.New(dao, passwordService)

	return &Services{
		AuthService:     authService,
		PasswordService: passwordService,
		UserService:     userService,
		TokenService:    tokenService,
	}
}
