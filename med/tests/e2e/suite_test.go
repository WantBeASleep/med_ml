//go:build e2e

package e2e_test

import (
	"fmt"
	"os"
	"testing"

	pb "med/internal/generated/grpc/service"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TestSuite struct {
	suite.Suite

	grpcClient pb.MedSrvClient
}

func (suite *TestSuite) SetupSuite() {
	conn, err := grpc.NewClient(
		os.Getenv("APP_URL"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(fmt.Sprintf("grpc connection failed: %v", err))
	}
	suite.grpcClient = pb.NewMedSrvClient(conn)
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
