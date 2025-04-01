//go:build e2e

package login_test

import (
	pb "auth/internal/generated/grpc/service"
	"auth/tests/e2e/flow"
	"fmt"

	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestLogin_Success() {
	data, err := flow.New(suite.deps, flow.RegisterUser).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	resp, err := suite.deps.Adapter.Login(suite.T().Context(), &pb.LoginIn{
		Email:    data.RegisterUser.Email,
		Password: data.RegisterUser.Password,
	})
	require.NoError(suite.T(), err)

	fmt.Println(resp)
	// accessToken, err := jwt.NewParser().ParseUnverified(resp.AccessToken, jwt.MapClaims{})
	// require.NoError(suite.T(), err)

	// refreshToken, err := jwt.NewParser().ParseUnverified(resp.RefreshToken, jwt.MapClaims{})
	// require.NoError(suite.T(), err)

}
