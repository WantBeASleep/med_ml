package tokens

import (
	"context"
	"errors"
	"net/http"

	"composition-api/internal/domain"
	auth_domain "composition-api/internal/domain/auth"
	api "composition-api/internal/generated/http/api"
)

func (h *handler) RefreshPost(ctx context.Context, req *api.RefreshPostReq) (api.RefreshPostRes, error) {
	accesstoken, refreshToken, err := h.services.TokensService.Refresh(ctx, auth_domain.Token(req.RefreshToken))
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrBadRequest):
			return &api.RefreshPostBadRequest{
				StatusCode: http.StatusBadRequest,
				Response: api.Error{
					Message: "Неверный формат запроса",
				},
			}, nil
		case errors.Is(err, domain.ErrUnauthorized):
			return &api.RefreshPostUnauthorized{
				StatusCode: http.StatusUnauthorized,
				Response: api.Error{
					Message: "Неверный или истекший refresh токен",
				},
			}, nil
		default:
			return nil, err
		}
	}

	return &api.RefreshPostOK{
		AccessToken:  accesstoken.String(),
		RefreshToken: refreshToken.String(),
	}, nil
}
