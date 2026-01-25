package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"composition-api/internal/server/billing"

	api "composition-api/internal/generated/http/api"
	"composition-api/internal/server/auth"
	"composition-api/internal/server/cytology"
	"composition-api/internal/server/download"
	"composition-api/internal/server/med"
	"composition-api/internal/server/register"
	"composition-api/internal/server/security"
	"composition-api/internal/server/tiler"
	"composition-api/internal/server/uzi"
	services "composition-api/internal/services"

	"github.com/ogen-go/ogen/ogenerrors"
)

type server struct {
	auth.AuthRoute
	uzi.UziRoute
	med.MedRoute
	register.RegisterRoute
	download.DownloadRoute
	billing.BillingRoute
	cytology.CytologyRoute
	tiler.TilerRoute
}

func New(services *services.Services) api.Handler {
	uziRoute := uzi.NewUziRoute(services)
	authRoute := auth.NewAuthRoute(services)
	medRoute := med.NewMedRoute(services)
	registerRoute := register.NewRegisterRoute(services)
	downloadRoute := download.NewDownloadRoute(services)
	billingRoute := billing.NewBillingRoute(services)
	cytologyRoute := cytology.NewCytologyRoute(services)
	tilerRoute := tiler.NewTilerRoute(services)

	return &server{
		UziRoute:      uziRoute,
		AuthRoute:     authRoute,
		MedRoute:      medRoute,
		RegisterRoute: registerRoute,
		DownloadRoute: downloadRoute,
		BillingRoute:  billingRoute,
		CytologyRoute: cytologyRoute,
		TilerRoute:    tilerRoute,
	}
}

func (s *server) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	// Проверяем, является ли ошибка ошибкой безопасности
	var securityErr *ogenerrors.SecurityError
	if errors.As(err, &securityErr) {
		// Проверяем, связана ли ошибка с токеном
		if errors.Is(err, security.ErrInvalidToken) || errors.Is(err, security.ErrUnauthorized) {
			return &api.ErrorStatusCode{
				StatusCode: http.StatusUnauthorized,
				Response: api.Error{
					Message: "Неверный или отсутствующий токен авторизации",
				},
			}
		}
		// Для других ошибок безопасности также возвращаем 401
		return &api.ErrorStatusCode{
			StatusCode: http.StatusUnauthorized,
			Response: api.Error{
				Message: "Ошибка авторизации",
			},
		}
	}

	// Проверяем, является ли ошибка напрямую ошибкой токена
	if errors.Is(err, security.ErrInvalidToken) || errors.Is(err, security.ErrUnauthorized) {
		return &api.ErrorStatusCode{
			StatusCode: http.StatusUnauthorized,
			Response: api.Error{
				Message: "Неверный или отсутствующий токен авторизации",
			},
		}
	}

	return &api.ErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: api.Error{
			Message: fmt.Sprint("Необработанная ошибка сервера: ", err.Error()),
		},
	}
}
