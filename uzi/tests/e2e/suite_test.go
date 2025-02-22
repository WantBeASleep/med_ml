//go:build e2e

// TODO: вынести в тестах общую часть в сьют
package e2e_test

import (
	"fmt"
	"os"
	"testing"

	pb "uzi/internal/generated/grpc/service"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type TestSuite struct {
	suite.Suite

	grpcClient pb.UziSrvClient
	dbusClient sarama.SyncProducer

	s3Client *minio.Client
	s3Bucket string
}

func (suite *TestSuite) SetupSuite() {
	conn, err := grpc.NewClient(
		os.Getenv("APP_URL"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(fmt.Sprintf("grpc connection failed: %v", err))
	}

	suite.dbusClient, err = sarama.NewSyncProducer(
		[]string{os.Getenv("BROKER_ADDRS")},
		nil,
	)
	if err != nil {
		panic(fmt.Sprintf("dbus connection failed: %v", err))
	}

	suite.s3Client, err = minio.New(os.Getenv("S3_ENDPOINT"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("S3_TOKEN_ACCESS"), os.Getenv("S3_TOKEN_SECRET"), ""),
		Secure: false,
	})
	if err != nil {
		panic(fmt.Sprintf("minio connection failed: %v", err))
	}

	// TODO: сделать прокидывание через env
	suite.s3Bucket = "uzi"
	suite.grpcClient = pb.NewUziSrvClient(conn)
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
