//go:build e2e

package e2e_test

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"testing"

	pb "auth/internal/generated/grpc/service"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TestSuite struct {
	suite.Suite

	client    pb.AuthSrvClient
	publicKey *rsa.PublicKey
}

func (suite *TestSuite) SetupSuite() {
	conn, err := grpc.NewClient(
		os.Getenv("APP_URL"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(fmt.Sprintf("grpc connection failed: %v", err))
	}

	publicBlock, _ := pem.Decode([]byte(os.Getenv("JWT_KEY_PUBLIC")))
	publicKey, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
	if err != nil {
		panic(fmt.Sprintf("parse public key: %v", err))
	}

	suite.publicKey = publicKey.(*rsa.PublicKey)
	suite.client = pb.NewAuthSrvClient(conn)
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
