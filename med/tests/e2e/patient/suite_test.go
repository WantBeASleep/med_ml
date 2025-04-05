//go:build e2e

package patient_test

import (
	"testing"

	e2e "med/tests/e2e"
	"med/tests/e2e/flow"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite

	deps *flow.Deps
}

func (suite *TestSuite) SetupSuite() {
	suite.deps = e2e.SetupDeps()
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
