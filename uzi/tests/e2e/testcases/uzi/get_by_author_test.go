//go:build e2e

package uzi_test

import (
	"github.com/stretchr/testify/require"

	pb "uzi/internal/generated/grpc/service"
	"uzi/tests/e2e/flow"
)

func (suite *TestSuite) TestGetUziByAuthor_Success() {
	data, err := flow.New(suite.deps, flow.DeviceInit, flow.UziInit).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	getResp, err := suite.deps.Adapter.GetUzisByAuthor(
		suite.T().Context(),
		&pb.GetUzisByAuthorIn{Author: data.Uzi.Author.String()},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 1, len(getResp.Uzis))
	require.Equal(suite.T(), data.Uzi.Id.String(), getResp.Uzis[0].Id)
	require.Equal(suite.T(), data.Uzi.Projection, getResp.Uzis[0].Projection)
	require.Equal(suite.T(), false, getResp.Uzis[0].Checked)
	require.Equal(suite.T(), data.Uzi.ExternalID.String(), getResp.Uzis[0].ExternalId)
	require.Equal(suite.T(), data.Uzi.Author.String(), getResp.Uzis[0].Author)
	require.Equal(suite.T(), pb.UziStatus_UZI_STATUS_NEW, getResp.Uzis[0].Status)
	require.Equal(suite.T(), int64(data.Uzi.DeviceID), getResp.Uzis[0].DeviceId)
}
