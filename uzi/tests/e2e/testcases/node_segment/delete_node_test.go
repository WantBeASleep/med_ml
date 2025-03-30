//go:build e2e

package node_segment_test

import (
	"math/rand"

	"github.com/stretchr/testify/require"

	pb "uzi/internal/generated/grpc/service"
	"uzi/tests/e2e/flow"
)

func (suite *TestSuite) TestDeleteNode_AiNodeDeleteNotAllowed() {
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
	require.Error(suite.T(), err)

	_, err = suite.deps.Adapter.GetSegmentsByNodeId(
		suite.T().Context(),
		&pb.GetSegmentsByNodeIdIn{NodeId: data.Nodes[0].Id.String()},
	)
	// TODO: после пра с обработкой ошибок, сделать нормальную проверку
	require.NoError(suite.T(), err)
}

func (suite *TestSuite) TestDeleteNode_Success() {
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
			Ai:       false,
			Node:     node,
			Segments: segments,
		},
	)
	require.NoError(suite.T(), err)

	_, err = suite.deps.Adapter.DeleteNode(
		suite.T().Context(),
		&pb.DeleteNodeIn{Id: createResp.NodeId},
	)
	require.NoError(suite.T(), err)

	_, err = suite.deps.Adapter.GetSegmentsByNodeId(
		suite.T().Context(),
		&pb.GetSegmentsByNodeIdIn{NodeId: createResp.NodeId},
	)
	// TODO: после пра с обработкой ошибок, сделать нормальную проверку
	require.Error(suite.T(), err)
}
