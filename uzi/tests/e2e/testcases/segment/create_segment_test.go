//go:build e2e

package segment_test

import (
	"math"
	"math/rand"

	"github.com/stretchr/testify/require"

	pb "uzi/internal/generated/grpc/service"
	"uzi/tests/e2e/flow"
)

func (suite *TestSuite) TestCreateSegment_AiNodeCreateNotAllowed() {
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

	_, err = suite.deps.Adapter.CreateSegment(
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
	require.Error(suite.T(), err)
}

func (suite *TestSuite) TestCreateSegment_Success() {
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

	createNodeResp, err := suite.deps.Adapter.CreateNodeWithSegments(
		suite.T().Context(),
		&pb.CreateNodeWithSegmentsIn{
			UziId:    data.Uzi.Id.String(),
			Node:     node,
			Segments: segments,
		},
	)
	require.NoError(suite.T(), err)

	contor := []byte(`[{"x": 1, "y": 1}]`)
	tirads23 := rand.Float64()
	tirads4 := rand.Float64()
	tirads5 := rand.Float64()

	createResp, err := suite.deps.Adapter.CreateSegment(
		suite.T().Context(),
		&pb.CreateSegmentIn{
			NodeId:    createNodeResp.NodeId,
			ImageId:   data.Images[1].Id.String(),
			Contor:    contor,
			Tirads_23: tirads23,
			Tirads_4:  tirads4,
			Tirads_5:  tirads5,
		},
	)
	require.NoError(suite.T(), err)

	getResp, err := suite.deps.Adapter.GetNodesWithSegmentsByImageId(
		suite.T().Context(),
		&pb.GetNodesWithSegmentsByImageIdIn{Id: data.Images[1].Id.String()},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 1, len(getResp.Segments))
	require.NotNil(suite.T(), getResp.Segments[0])
	require.Equal(suite.T(), createResp.Id, getResp.Segments[0].Id)
	require.Equal(suite.T(), createNodeResp.NodeId, getResp.Segments[0].NodeId)
	require.Equal(suite.T(), data.Images[1].Id.String(), getResp.Segments[0].ImageId)
	require.Equal(suite.T(), contor, getResp.Segments[0].Contor)
	require.False(suite.T(), getResp.Segments[0].Ai)
	require.True(suite.T(), math.Abs(tirads23-getResp.Segments[0].Tirads_23) < 0.0001)
	require.True(suite.T(), math.Abs(tirads4-getResp.Segments[0].Tirads_4) < 0.0001)
	require.True(suite.T(), math.Abs(tirads5-getResp.Segments[0].Tirads_5) < 0.0001)
}
