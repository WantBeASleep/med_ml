//go:build e2e

package device_test

import (
	"testing"

	"uzi/tests/e2e/flow"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite

	deps *flow.Deps
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
