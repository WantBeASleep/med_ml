package tokens

import (
	"context"

	api "composition-api/internal/generated/http/api"
)

func (h *handler) LoginPost(ctx context.Context, req *api.LoginPostReq) (api.LoginPostRes, error) {
	accesstoken, refreshToken, err := h.services.TokensService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &api.LoginPostOK{
		AccessToken:  accesstoken.String(),
		RefreshToken: refreshToken.String(),
	}, nil
}
