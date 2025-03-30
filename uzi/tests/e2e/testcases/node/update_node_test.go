//go:build e2e

package node_test

import (
	"fmt"
	"math/rand"

	"github.com/stretchr/testify/require"

	pb "uzi/internal/generated/grpc/service"
	"uzi/tests/e2e/flow"
)

func (suite *TestSuite) TestUpdateNode_AIUpdateNotAllowed() {
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

	fmt.Println(data.Nodes[0].Id.String(), "это меняем")

	_, err = suite.deps.Adapter.UpdateNode(
		suite.T().Context(),
		&pb.UpdateNodeIn{
			Id:        data.Nodes[0].Id.String(),
			Tirads_23: &tirads23,
			Tirads_4:  &tirads4,
			Tirads_5:  &tirads5,
		},
	)
	require.Error(suite.T(), err)
}

func (suite *TestSuite) TestUpdateNode_Success() {
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

	tirads23 := rand.Float64()
	tirads4 := rand.Float64()
	tirads5 := rand.Float64()

	updateResp, err := suite.deps.Adapter.UpdateNode(
		suite.T().Context(),
		&pb.UpdateNodeIn{
			Id:        createResp.NodeId,
			Tirads_23: &tirads23,
			Tirads_4:  &tirads4,
			Tirads_5:  &tirads5,
		},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), tirads23, updateResp.Node.Tirads_23)
	require.Equal(suite.T(), tirads4, updateResp.Node.Tirads_4)
	require.Equal(suite.T(), tirads5, updateResp.Node.Tirads_5)
}
