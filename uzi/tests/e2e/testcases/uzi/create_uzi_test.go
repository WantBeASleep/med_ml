//go:build e2e

package uzi_test

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"uzi/internal/domain"

	pb "uzi/internal/generated/grpc/service"
	"uzi/internal/server/mappers"
	"uzi/tests/e2e/flow"
)

func (suite *TestSuite) TestCreateUzi_Success() {
	data, err := flow.New(suite.deps, flow.DeviceInit).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	projection := domain.UziProjectionLong
	externalID := uuid.New()
	author := uuid.New()
	description := gofakeit.Word()
	createResp, err := suite.deps.Adapter.CreateUzi(
		suite.T().Context(),
		&pb.CreateUziIn{
			DeviceId:    int64(data.Device.Id),
			Projection:  mappers.UziProjectionMap[projection],
			ExternalId:  externalID.String(),
			Author:      author.String(),
			Description: &description,
		},
	)
	require.NoError(suite.T(), err)

	getResp, err := suite.deps.Adapter.GetUziById(
		suite.T().Context(),
		&pb.GetUziByIdIn{Id: createResp.Id},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), createResp.Id, getResp.Uzi.Id)
	require.Equal(suite.T(), projection, mappers.UziProjectionReverseMap[getResp.Uzi.Projection])
	require.Equal(suite.T(), false, getResp.Uzi.Checked)
	require.Equal(suite.T(), externalID.String(), getResp.Uzi.ExternalId)
	require.Equal(suite.T(), author.String(), getResp.Uzi.Author)
	require.Equal(suite.T(), pb.UziStatus_UZI_STATUS_NEW, getResp.Uzi.Status)
	require.Equal(suite.T(), int64(data.Device.Id), getResp.Uzi.DeviceId)
	require.Equal(suite.T(), description, *getResp.Uzi.Description)
}
