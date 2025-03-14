//go:build e2e

package node_test

import (
	"math/rand"

	"github.com/stretchr/testify/require"

	pb "uzi/internal/generated/grpc/service"
	"uzi/tests/e2e/flow"
)

func (suite *TestSuite) TestUpdateNode_Success() {
	data, err := flow.New(
		suite.deps,
		flow.DeviceInit,
		flow.UziInit,
		flow.TiffSplit,
		flow.SaveNodesWithSegments,
	).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	tirads23 := rand.Float64()
	tirads4 := rand.Float64()
	tirads5 := rand.Float64()

	updateResp, err := suite.deps.Adapter.UpdateNode(
		suite.T().Context(),
		&pb.UpdateNodeIn{
			Id:        data.Nodes[0].Id.String(),
			Tirads_23: &tirads23,
			Tirads_4:  &tirads4,
			Tirads_5:  &tirads5,
		},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), data.Nodes[0].Id.String(), updateResp.Node.Id)
	require.Equal(suite.T(), data.Uzi.Id.String(), updateResp.Node.UziId)
	require.Equal(suite.T(), tirads23, updateResp.Node.Tirads_23)
	require.Equal(suite.T(), tirads4, updateResp.Node.Tirads_4)
	require.Equal(suite.T(), tirads5, updateResp.Node.Tirads_5)
}
