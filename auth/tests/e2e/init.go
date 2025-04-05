//go:build e2e

package e2e_test

import (
	"fmt"
	"os"

	pb "auth/internal/generated/grpc/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"auth/tests/e2e/flow"
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

	deps.Adapter = pb.NewAuthSrvClient(conn)

	return deps
}
