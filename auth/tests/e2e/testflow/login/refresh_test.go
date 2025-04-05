//go:build e2e

package login_test

import (
	"time"

	pb "auth/internal/generated/grpc/service"
	"auth/tests/e2e/flow"

	"auth/tests/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestRefresh_Success() {
	data, err := flow.New(suite.deps, flow.RegisterUser, flow.Login).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	resp, err := suite.deps.Adapter.Refresh(suite.T().Context(), &pb.RefreshIn{
		RefreshToken: data.Tokens.Refresh.String(),
	})
	require.NoError(suite.T(), err)

	// access token
	accessToken, _, err := jwt.NewParser().ParseUnverified(resp.AccessToken, jwt.MapClaims{})
	require.NoError(suite.T(), err)
	expt, err := accessToken.Claims.GetExpirationTime()
	require.NoError(suite.T(), err)
	require.True(suite.T(), expt.After(time.Now()))

	claims := accessToken.Claims.(jwt.MapClaims)
	require.Equal(suite.T(), data.RegisterUser.Id.String(), utils.MustParseFromClaims(suite.T(), "id", claims))
	require.Equal(suite.T(), data.RegisterUser.Role.String(), utils.MustParseFromClaims(suite.T(), "role", claims))

	// refresh token
	refreshToken, _, err := jwt.NewParser().ParseUnverified(resp.RefreshToken, jwt.MapClaims{})
	require.NoError(suite.T(), err)
	expt, err = refreshToken.Claims.GetExpirationTime()
	require.NoError(suite.T(), err)
	require.True(suite.T(), expt.After(time.Now()))

	claims = refreshToken.Claims.(jwt.MapClaims)
	require.Equal(suite.T(), data.RegisterUser.Id.String(), utils.MustParseFromClaims(suite.T(), "id", claims))
	require.Equal(suite.T(), data.RegisterUser.Role.String(), utils.MustParseFromClaims(suite.T(), "role", claims))
}
