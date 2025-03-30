//go:build e2e

package node_test

import (
	"github.com/stretchr/testify/require"

	pb "uzi/internal/generated/grpc/service"
	"uzi/tests/e2e/flow"
)

func (suite *TestSuite) TestGetNodesByUziId_Success() {
	data, err := flow.New(
		suite.deps,
		flow.DeviceInit,
		flow.UziInit,
		flow.TiffSplit,
		flow.SaveNodesWithSegments,
	).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	getResp, err := suite.deps.Adapter.GetNodesByUziId(
		suite.T().Context(),
		&pb.GetNodesByUziIdIn{UziId: data.Uzi.Id.String()},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), len(data.Nodes), len(getResp.Nodes))
	for i, node := range getResp.Nodes {
		require.Equal(suite.T(), data.Nodes[i].Id.String(), node.Id)
		require.Equal(suite.T(), data.Nodes[i].Ai, node.Ai)
		require.Equal(suite.T(), pb.NodeValidation_NODE_VALIDATION_NULL, *node.Validation)
		require.Equal(suite.T(), data.Uzi.Id.String(), node.UziId)
		require.Equal(suite.T(), data.Nodes[i].Tirads23, node.Tirads_23)
		require.Equal(suite.T(), data.Nodes[i].Tirads4, node.Tirads_4)
		require.Equal(suite.T(), data.Nodes[i].Tirads5, node.Tirads_5)
		require.Nil(suite.T(), node.Description)
	}
}
