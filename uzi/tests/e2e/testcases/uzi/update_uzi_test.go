//go:build e2e

package uzi_test

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"

	pb "uzi/internal/generated/grpc/service"
	"uzi/tests/e2e/flow"
)

func (suite *TestSuite) TestUpdateUzi_Success() {
	data, err := flow.New(suite.deps, flow.DeviceInit, flow.UziInit).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	projection := gofakeit.Word()
	checked := true

	resp, err := suite.deps.Adapter.UpdateUzi(
		suite.T().Context(),
		&pb.UpdateUziIn{
			Id:         data.Uzi.Id.String(),
			Projection: &projection,
			Checked:    &checked,
		},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), data.Uzi.Id.String(), resp.Uzi.Id)
	require.Equal(suite.T(), projection, resp.Uzi.Projection)
	require.Equal(suite.T(), checked, resp.Uzi.Checked)
	require.Equal(suite.T(), data.Uzi.ExternalID.String(), resp.Uzi.ExternalId)
	require.Equal(suite.T(), data.Uzi.Author.String(), resp.Uzi.Author)
	require.Equal(suite.T(), pb.UziStatus_UZI_STATUS_NEW, resp.Uzi.Status)
	require.Equal(suite.T(), int64(data.Uzi.DeviceID), resp.Uzi.DeviceId)
	require.NotNil(suite.T(), resp.Uzi.Description)
	require.Equal(suite.T(), *data.Uzi.Description, *resp.Uzi.Description)
}
