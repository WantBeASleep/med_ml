//go:build e2e

package node_segment_test

import (
	"github.com/stretchr/testify/require"

	pb "uzi/internal/generated/grpc/service"
	"uzi/tests/e2e/flow"
)

func (suite *TestSuite) TestDeleteSegment_Success() {
	data, err := flow.New(
		suite.deps,
		flow.DeviceInit,
		flow.UziInit,
		flow.TiffSplit,
		flow.SaveNodesWithSegments,
	).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	_, err = suite.deps.Adapter.DeleteSegment(
		suite.T().Context(),
		&pb.DeleteSegmentIn{Id: data.Segments[0].Id.String()},
	)
	require.NoError(suite.T(), err)

	resp, err := suite.deps.Adapter.GetSegmentsByNodeId(
		suite.T().Context(),
		&pb.GetSegmentsByNodeIdIn{NodeId: data.Segments[0].NodeID.String()},
	)
	require.NoError(suite.T(), err)

	find := false
	for _, v := range resp.Segments {
		if v.Id == data.Segments[0].Id.String() {
			find = true
			break
		}
	}
	require.False(suite.T(), find)
}
