//go:build e2e

package segment_test

import (
	"github.com/stretchr/testify/require"

	"uzi/internal/domain"
	pb "uzi/internal/generated/grpc/service"
	"uzi/tests/e2e/flow"
)

func (suite *TestSuite) TestGetSegmentsByNodeId_Success() {
	data, err := flow.New(
		suite.deps,
		flow.DeviceInit,
		flow.UziInit,
		flow.TiffSplit,
		flow.SaveNodesWithSegments,
	).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	getResp, err := suite.deps.Adapter.GetSegmentsByNodeId(
		suite.T().Context(),
		&pb.GetSegmentsByNodeIdIn{NodeId: data.Nodes[0].Id.String()},
	)
	require.NoError(suite.T(), err)

	expectedSegments := map[string]domain.Segment{}
	for _, segment := range data.Segments {
		if segment.NodeID == data.Nodes[0].Id {
			expectedSegments[segment.Id.String()] = segment
		}
	}
	require.Equal(suite.T(), len(expectedSegments), len(getResp.Segments))
	for _, segment := range getResp.Segments {
		require.Equal(suite.T(), expectedSegments[segment.Id].Id.String(), segment.Id)
		require.Equal(suite.T(), expectedSegments[segment.Id].NodeID.String(), segment.NodeId)
		require.Equal(suite.T(), expectedSegments[segment.Id].ImageID.String(), segment.ImageId)
		require.Equal(suite.T(), string(expectedSegments[segment.Id].Contor), string(segment.Contor))
		require.Equal(suite.T(), expectedSegments[segment.Id].Tirads23, segment.Tirads_23)
		require.Equal(suite.T(), expectedSegments[segment.Id].Tirads4, segment.Tirads_4)
		require.Equal(suite.T(), expectedSegments[segment.Id].Tirads5, segment.Tirads_5)
	}
}
