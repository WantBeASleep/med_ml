//go:build e2e

package e2e_test

import (
	"fmt"
	"os"

	pb "uzi/internal/generated/grpc/service"

	"github.com/IBM/sarama"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"uzi/tests/e2e/flow"

	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func SetupDeps() *flow.Deps {
	deps := &flow.Deps{}

	conn, err := grpc.NewClient(
		os.Getenv("APP_URL"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(fmt.Sprintf("grpc connection failed: %v", err))
	}

	deps.Adapter = pb.NewUziSrvClient(conn)

	deps.Dbus, err = sarama.NewSyncProducer(
		[]string{os.Getenv("BROKER_ADDRS")},
		nil,
	)
	if err != nil {
		panic(fmt.Sprintf("dbus connection failed: %v", err))
	}

	deps.S3, err = minio.New(os.Getenv("S3_ENDPOINT"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("S3_TOKEN_ACCESS"), os.Getenv("S3_TOKEN_SECRET"), ""),
		Secure: false,
	})
	if err != nil {
		panic(fmt.Sprintf("minio connection failed: %v", err))
	}

	// TODO: вбить как константу в репозитории
	deps.Bucket = "uzi"

	return deps
}
