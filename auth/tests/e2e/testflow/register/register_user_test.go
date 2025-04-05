//go:build e2e

package login_test

import (
	pb "auth/internal/generated/grpc/service"
	"auth/internal/server/mappers"
	"auth/tests/e2e/flow"
	"auth/tests/utils"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestRegisterUser_Success() {
	email := gofakeit.Email()
	password := gofakeit.Password(true, true, true, true, false, 10)

	_, err := suite.deps.Adapter.RegisterUser(suite.T().Context(), &pb.RegisterUserIn{
		Email:    email,
		Password: password,
		Role:     pb.Role_ROLE_DOCTOR,
	})
	require.NoError(suite.T(), err)

	resp, err := suite.deps.Adapter.Login(suite.T().Context(), &pb.LoginIn{
		Email:    email,
		Password: password,
	})
	require.NoError(suite.T(), err)

	accessToken, _, err := jwt.NewParser().ParseUnverified(resp.AccessToken, jwt.MapClaims{})
	require.NoError(suite.T(), err)
	claims := accessToken.Claims.(jwt.MapClaims)
	require.Equal(suite.T(), mappers.RoleReversedMap[pb.Role_ROLE_DOCTOR].String(), utils.MustParseFromClaims(suite.T(), "role", claims))
}

func (suite *TestSuite) TestRegisterUser_UserAlreadyExists() {
	data, err := flow.New(suite.deps, flow.RegisterUser).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	_, err = suite.deps.Adapter.RegisterUser(suite.T().Context(), &pb.RegisterUserIn{
		Email:    data.RegisterUser.Email,
		Password: data.RegisterUser.Password,
		Role:     pb.Role_ROLE_PATIENT,
	})
	require.Error(suite.T(), err)
}

func (suite *TestSuite) TestRegisterUser_UnRegisteredUser() {
	data, err := flow.New(suite.deps, flow.CreateUnRegisterUser).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	pass := "qwerty123"

	_, err = suite.deps.Adapter.RegisterUser(suite.T().Context(), &pb.RegisterUserIn{
		Email:    data.UnRegisterUser.Email,
		Password: pass,
		Role:     pb.Role_ROLE_PATIENT,
	})
	require.NoError(suite.T(), err)

	_, err = suite.deps.Adapter.Login(suite.T().Context(), &pb.LoginIn{
		Email:    data.UnRegisterUser.Email,
		Password: pass,
	})
	require.NoError(suite.T(), err)
}

func (suite *TestSuite) TestRegisterUser_ReUnregisterUser() {
	data, err := flow.New(suite.deps, flow.CreateUnRegisterUser).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	_, err = suite.deps.Adapter.CreateUnRegisteredUser(suite.T().Context(), &pb.CreateUnRegisteredUserIn{
		Email: data.UnRegisterUser.Email,
	})
	require.Error(suite.T(), err)
}

func (suite *TestSuite) TestRegisterUser_ReUnregisterUserOnRegisteredUser() {
	data, err := flow.New(suite.deps, flow.RegisterUser).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	_, err = suite.deps.Adapter.CreateUnRegisteredUser(suite.T().Context(), &pb.CreateUnRegisteredUserIn{
		Email: data.RegisterUser.Email,
	})
	require.Error(suite.T(), err)
}
