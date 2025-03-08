//go:build e2e

package segment_test

import (
	"math"
	"math/rand"

	"github.com/stretchr/testify/require"

	pb "uzi/internal/generated/grpc/service"
	"uzi/tests/e2e/flow"
)

func (suite *TestSuite) TestUpdateSegment_Success() {
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

	updateResp, err := suite.deps.Adapter.UpdateSegment(
		suite.T().Context(),
		&pb.UpdateSegmentIn{
			Id:        data.Segments[0].Id.String(),
			Tirads_23: &tirads23,
			Tirads_4:  &tirads4,
			Tirads_5:  &tirads5,
		},
	)
	require.NoError(suite.T(), err)

	require.Equal(suite.T(), data.Segments[0].Id.String(), updateResp.Segment.Id)
	require.Equal(suite.T(), data.Segments[0].NodeID.String(), updateResp.Segment.NodeId)
	require.Equal(suite.T(), data.Segments[0].ImageID.String(), updateResp.Segment.ImageId)
	require.Equal(suite.T(), string(data.Segments[0].Contor), string(updateResp.Segment.Contor))
	require.True(suite.T(), math.Abs(tirads23-updateResp.Segment.Tirads_23) < 0.0001)
	require.True(suite.T(), math.Abs(tirads4-updateResp.Segment.Tirads_4) < 0.0001)
	require.True(suite.T(), math.Abs(tirads5-updateResp.Segment.Tirads_5) < 0.0001)
}
