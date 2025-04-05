//go:build e2e

package e2e_test

import (
	"fmt"
	"os"

	pb "med/internal/generated/grpc/service"
	"med/tests/e2e/flow"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	deps.Adapter = pb.NewMedSrvClient(conn)

	return deps
}
