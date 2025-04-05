//go:build e2e

package image_test

import (
	"github.com/stretchr/testify/require"

	pb "uzi/internal/generated/grpc/service"
	"uzi/tests/e2e/flow"
)

func (suite *TestSuite) TestGetImagesByUziId_Success() {
	data, err := flow.New(suite.deps, flow.DeviceInit, flow.UziInit, flow.TiffSplit).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	getResp, err := suite.deps.Adapter.GetImagesByUziId(
		suite.T().Context(),
		&pb.GetImagesByUziIdIn{UziId: data.Uzi.Id.String()},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), len(data.Images), len(getResp.Images))
	for i, image := range getResp.Images {
		require.Equal(suite.T(), data.Images[i].Id.String(), image.Id)
		require.Equal(suite.T(), data.Uzi.Id.String(), image.UziId)
		require.Equal(suite.T(), int64(data.Images[i].Page), image.Page)
	}
}
