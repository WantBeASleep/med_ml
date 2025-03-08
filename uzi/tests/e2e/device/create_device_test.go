//go:build e2e

package device_test

import (
	"context"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "uzi/internal/generated/grpc/service"
)

func (suite *TestSuite) TestCreateDevice_Success() {
	deviceNames := make([]string, 0, 5)
	for range 5 {
		deviceNames = append(deviceNames, gofakeit.School())
	}

	for _, deviceName := range deviceNames {
		_, err := suite.deps.Adapter.CreateDevice(context.Background(), &pb.CreateDeviceIn{Name: deviceName})
		require.NoError(suite.T(), err)
	}

	resp, err := suite.deps.Adapter.GetDeviceList(context.Background(), &emptypb.Empty{})
	require.NoError(suite.T(), err)

	gotDevices := make(map[string]struct{})
	for _, device := range resp.Devices {
		gotDevices[device.Name] = struct{}{}
	}

	for _, deviceName := range deviceNames {
		_, ok := gotDevices[deviceName]
		require.True(suite.T(), ok)
	}
}
