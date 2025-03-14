//go:build e2e

package segment_test

import (
	"math"
	"math/rand"

	"github.com/stretchr/testify/require"

	pb "uzi/internal/generated/grpc/service"
	"uzi/tests/e2e/flow"
)

func (suite *TestSuite) TestCreateSegment_Success() {
	data, err := flow.New(
		suite.deps,
		flow.DeviceInit,
		flow.UziInit,
		flow.TiffSplit,
		flow.SaveNodesWithSegments,
	).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	contor := []byte(`[{"x": 1, "y": 1}]`)
	tirads23 := rand.Float64()
	tirads4 := rand.Float64()
	tirads5 := rand.Float64()

	createResp, err := suite.deps.Adapter.CreateSegment(
		suite.T().Context(),
		&pb.CreateSegmentIn{
			NodeId:    data.Nodes[0].Id.String(),
			ImageId:   data.Images[0].Id.String(),
			Contor:    contor,
			Tirads_23: tirads23,
			Tirads_4:  tirads4,
			Tirads_5:  tirads5,
		},
	)
	require.NoError(suite.T(), err)

	getResp, err := suite.deps.Adapter.GetSegmentsByNodeId(
		suite.T().Context(),
		&pb.GetSegmentsByNodeIdIn{NodeId: data.Nodes[0].Id.String()},
	)
	require.NoError(suite.T(), err)

	var findedSegment *pb.Segment
	for _, segment := range getResp.Segments {
		if segment.Id == createResp.Id {
			findedSegment = segment
			break
		}
	}
	require.NotNil(suite.T(), findedSegment)
	require.Equal(suite.T(), createResp.Id, findedSegment.Id)
	require.Equal(suite.T(), data.Nodes[0].Id.String(), findedSegment.NodeId)
	require.Equal(suite.T(), data.Images[0].Id.String(), findedSegment.ImageId)
	require.Equal(suite.T(), contor, findedSegment.Contor)
	require.True(suite.T(), math.Abs(tirads23-findedSegment.Tirads_23) < 0.0001)
	require.True(suite.T(), math.Abs(tirads4-findedSegment.Tirads_4) < 0.0001)
	require.True(suite.T(), math.Abs(tirads5-findedSegment.Tirads_5) < 0.0001)
}
