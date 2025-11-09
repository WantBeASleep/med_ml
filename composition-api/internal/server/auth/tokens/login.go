package tokens

import (
	"context"
	"errors"
	"net/http"

	"composition-api/internal/domain"
	api "composition-api/internal/generated/http/api"
)

func (h *handler) LoginPost(ctx context.Context, req *api.LoginPostReq) (api.LoginPostRes, error) {
	accesstoken, refreshToken, err := h.services.TokensService.Login(ctx, req.Email, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrBadRequest):
			return &api.LoginPostBadRequest{
				StatusCode: http.StatusBadRequest,
				Response: api.Error{
					Message: "Неверный формат запроса",
				},
			}, nil
		case errors.Is(err, domain.ErrUnauthorized):
			return &api.LoginPostUnauthorized{
				StatusCode: http.StatusUnauthorized,
				Response: api.Error{
					Message: "Неверный email или пароль",
				},
			}, nil
		default:
			return nil, err
		}
	}

	return &api.LoginPostOK{
		AccessToken:  accesstoken.String(),
		RefreshToken: refreshToken.String(),
	}, nil
}
