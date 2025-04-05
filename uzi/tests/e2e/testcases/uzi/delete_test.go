//go:build e2e

package uzi_test

import (
	"github.com/stretchr/testify/require"

	pb "uzi/internal/generated/grpc/service"
	"uzi/tests/e2e/flow"
)

func (suite *TestSuite) TestDeleteUzi_Success() {
	data, err := flow.New(suite.deps, flow.DeviceInit, flow.UziInit).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	_, err = suite.deps.Adapter.DeleteUzi(
		suite.T().Context(),
		&pb.DeleteUziIn{Id: data.Uzi.Id.String()},
	)
	require.NoError(suite.T(), err)

	_, err = suite.deps.Adapter.GetUziById(
		suite.T().Context(),
		&pb.GetUziByIdIn{Id: data.Uzi.Id.String()},
	)
	require.Error(suite.T(), err)
}
