//go:build e2e

package node_segment_test

import (
	"math"
	"math/rand"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"

	pb "uzi/internal/generated/grpc/service"
	"uzi/tests/e2e/flow"
)

func (suite *TestSuite) TestCreateNodeWithSegment_Success() {
	data, err := flow.New(
		suite.deps,
		flow.DeviceInit,
		flow.UziInit,
		flow.TiffSplit,
	).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	description := gofakeit.Word()

	node := &pb.CreateNodeWithSegmentsIn_Node{
		Tirads_23:   rand.Float64(),
		Tirads_4:    rand.Float64(),
		Tirads_5:    rand.Float64(),
		Description: &description,
	}

	segments := []*pb.CreateNodeWithSegmentsIn_Segment{
		{
			ImageId:   data.Images[0].Id.String(),
			Contor:    []byte(`[{"x": 1, "y": 1}]`),
			Tirads_23: rand.Float64(),
			Tirads_4:  rand.Float64(),
			Tirads_5:  rand.Float64(),
		},
		{
			ImageId:   data.Images[1].Id.String(),
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

	getResp, err := suite.deps.Adapter.GetNodesWithSegmentsByImageId(
		suite.T().Context(),
		&pb.GetNodesWithSegmentsByImageIdIn{Id: data.Images[0].Id.String()},
	)
	require.NoError(suite.T(), err)

	require.Equal(suite.T(), len(getResp.Nodes), 1)
	require.Equal(suite.T(), getResp.Nodes[0].Id, createResp.NodeId)
	require.Equal(suite.T(), getResp.Nodes[0].Ai, false)
	require.Nil(suite.T(), getResp.Nodes[0].Validation)
	require.Equal(suite.T(), getResp.Nodes[0].UziId, data.Uzi.Id.String())
	require.True(suite.T(), math.Abs(getResp.Nodes[0].Tirads_23-node.Tirads_23) < 0.0001)
	require.True(suite.T(), math.Abs(getResp.Nodes[0].Tirads_4-node.Tirads_4) < 0.0001)
	require.True(suite.T(), math.Abs(getResp.Nodes[0].Tirads_5-node.Tirads_5) < 0.0001)
	require.NotNil(suite.T(), getResp.Nodes[0].Description)
	require.Equal(suite.T(), description, *getResp.Nodes[0].Description)

	require.Equal(suite.T(), len(getResp.Segments), 1)
	require.Equal(suite.T(), getResp.Segments[0].Id, createResp.SegmentIds[0])
	require.Equal(suite.T(), getResp.Segments[0].ImageId, segments[0].ImageId)
	require.Equal(suite.T(), getResp.Segments[0].Contor, segments[0].Contor)
	require.Equal(suite.T(), getResp.Segments[0].Ai, false)
	require.True(suite.T(), math.Abs(getResp.Segments[0].Tirads_23-segments[0].Tirads_23) < 0.0001)
	require.True(suite.T(), math.Abs(getResp.Segments[0].Tirads_4-segments[0].Tirads_4) < 0.0001)
	require.True(suite.T(), math.Abs(getResp.Segments[0].Tirads_5-segments[0].Tirads_5) < 0.0001)

	getResp, err = suite.deps.Adapter.GetNodesWithSegmentsByImageId(
		suite.T().Context(),
		&pb.GetNodesWithSegmentsByImageIdIn{Id: data.Images[1].Id.String()},
	)
	require.NoError(suite.T(), err)

	require.Equal(suite.T(), len(getResp.Segments), 1)
	require.Equal(suite.T(), getResp.Segments[0].Id, createResp.SegmentIds[1])
	require.Equal(suite.T(), getResp.Segments[0].ImageId, segments[1].ImageId)
	require.Equal(suite.T(), getResp.Segments[0].Contor, segments[1].Contor)
	require.Equal(suite.T(), getResp.Segments[0].Ai, false)
	require.True(suite.T(), math.Abs(getResp.Segments[0].Tirads_23-segments[1].Tirads_23) < 0.0001)
	require.True(suite.T(), math.Abs(getResp.Segments[0].Tirads_4-segments[1].Tirads_4) < 0.0001)
	require.True(suite.T(), math.Abs(getResp.Segments[0].Tirads_5-segments[1].Tirads_5) < 0.0001)
}
