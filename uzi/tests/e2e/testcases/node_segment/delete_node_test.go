//go:build e2e

package node_segment_test

import (
	"github.com/stretchr/testify/require"

	pb "uzi/internal/generated/grpc/service"
	"uzi/tests/e2e/flow"
)

func (suite *TestSuite) TestDeleteNode_Success() {
	data, err := flow.New(
		suite.deps,
		flow.DeviceInit,
		flow.UziInit,
		flow.TiffSplit,
		flow.SaveNodesWithSegments,
	).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	_, err = suite.deps.Adapter.DeleteNode(
		suite.T().Context(),
		&pb.DeleteNodeIn{Id: data.Nodes[0].Id.String()},
	)
	require.NoError(suite.T(), err)

	_, err = suite.deps.Adapter.GetSegmentsByNodeId(
		suite.T().Context(),
		&pb.GetSegmentsByNodeIdIn{NodeId: data.Nodes[0].Id.String()},
	)
	// TODO: после пра с обработкой ошибок, сделать нормальную проверку
	require.Error(suite.T(), err)
}
