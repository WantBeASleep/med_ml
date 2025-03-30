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
	).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	node := &pb.CreateNodeWithSegmentsIn_Node{
		Tirads_23: rand.Float64(),
		Tirads_4:  rand.Float64(),
		Tirads_5:  rand.Float64(),
	}

	segments := []*pb.CreateNodeWithSegmentsIn_Segment{
		{
			ImageId:   data.Images[0].Id.String(),
			Contor:    []byte(`[{"x": 1, "y": 1}]`),
			Tirads_23: rand.Float64(),
			Tirads_4:  rand.Float64(),
			Tirads_5:  rand.Float64(),
		},
	}

	createResp, err := suite.deps.Adapter.CreateNodeWithSegments(
		suite.T().Context(),
		&pb.CreateNodeWithSegmentsIn{
			UziId:    data.Uzi.Id.String(),
			Node:     node,
			Segments: segments,
		},
	)
	require.NoError(suite.T(), err)

	tirads23 := rand.Float64()
	tirads4 := rand.Float64()
	tirads5 := rand.Float64()

	updateResp, err := suite.deps.Adapter.UpdateSegment(
		suite.T().Context(),
		&pb.UpdateSegmentIn{
			Id:        createResp.SegmentIds[0],
			Tirads_23: &tirads23,
			Tirads_4:  &tirads4,
			Tirads_5:  &tirads5,
		},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), createResp.SegmentIds[0], updateResp.Segment.Id)
	require.Equal(suite.T(), createResp.NodeId, updateResp.Segment.NodeId)
	require.Equal(suite.T(), data.Images[0].Id.String(), updateResp.Segment.ImageId)
	require.Equal(suite.T(), string(segments[0].Contor), string(updateResp.Segment.Contor))
	require.True(suite.T(), math.Abs(tirads23-updateResp.Segment.Tirads_23) < 0.0001)
	require.True(suite.T(), math.Abs(tirads4-updateResp.Segment.Tirads_4) < 0.0001)
	require.True(suite.T(), math.Abs(tirads5-updateResp.Segment.Tirads_5) < 0.0001)

}

func (suite *TestSuite) TestUpdateSegment_AIUpdateNotAllowed() {
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

	_, err = suite.deps.Adapter.UpdateSegment(
		suite.T().Context(),
		&pb.UpdateSegmentIn{
			Id:        data.Segments[0].Id.String(),
			Tirads_23: &tirads23,
			Tirads_4:  &tirads4,
			Tirads_5:  &tirads5,
		},
	)
	require.Error(suite.T(), err)
}
