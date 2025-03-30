//go:build e2e

package uzi_test

import (
	"github.com/stretchr/testify/require"

	pb "uzi/internal/generated/grpc/service"
	"uzi/tests/e2e/flow"
)

func (suite *TestSuite) TestGetUziById_Success() {
	data, err := flow.New(suite.deps, flow.DeviceInit, flow.UziInit).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	getResp, err := suite.deps.Adapter.GetUziById(
		suite.T().Context(),
		&pb.GetUziByIdIn{Id: data.Uzi.Id.String()},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), data.Uzi.Id.String(), getResp.Uzi.Id)
	require.Equal(suite.T(), data.Uzi.Projection, getResp.Uzi.Projection)
	require.Equal(suite.T(), false, getResp.Uzi.Checked)
	require.Equal(suite.T(), data.Uzi.ExternalID.String(), getResp.Uzi.ExternalId)
	require.Equal(suite.T(), data.Uzi.Author.String(), getResp.Uzi.Author)
	require.Equal(suite.T(), pb.UziStatus_UZI_STATUS_NEW, getResp.Uzi.Status)
	require.Equal(suite.T(), int64(data.Uzi.DeviceID), getResp.Uzi.DeviceId)
	require.NotNil(suite.T(), getResp.Uzi.Description)
	require.Equal(suite.T(), *data.Uzi.Description, *getResp.Uzi.Description)
}
