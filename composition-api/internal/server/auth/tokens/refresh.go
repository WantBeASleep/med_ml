package tokens

import (
	"context"

	domain "composition-api/internal/domain/auth"
	api "composition-api/internal/generated/http/api"
)

func (h *handler) RefreshPost(ctx context.Context, req *api.RefreshPostReq) (api.RefreshPostRes, error) {
	accesstoken, refreshToken, err := h.services.TokensService.Refresh(ctx, domain.Token(req.RefreshToken))
	if err != nil {
		return nil, err
	}

	return &api.RefreshPostOK{
		AccessToken:  accesstoken.String(),
		RefreshToken: refreshToken.String(),
	}, nil
}
