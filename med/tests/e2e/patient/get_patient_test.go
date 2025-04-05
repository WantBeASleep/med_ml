//go:build e2e

package patient_test

import (
	"time"

	pb "med/internal/generated/grpc/service"
	"med/tests/e2e/flow"

	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestGetPatient_Success() {
	data, err := flow.New(suite.deps, flow.CreatePatient).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	getResp, err := suite.deps.Adapter.GetPatient(
		suite.T().Context(),
		&pb.GetPatientIn{Id: data.Patient.Id.String()},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), data.Patient.Id.String(), getResp.Patient.Id)
	require.Equal(suite.T(), data.Patient.FullName, getResp.Patient.Fullname)
	require.Equal(suite.T(), data.Patient.Email, getResp.Patient.Email)
	require.Equal(suite.T(), data.Patient.Policy, getResp.Patient.Policy)
	require.Equal(suite.T(), data.Patient.Active, getResp.Patient.Active)
	require.Equal(suite.T(), data.Patient.Malignancy, getResp.Patient.Malignancy)
	require.Equal(suite.T(), data.Patient.BirthDate.Format(time.RFC3339), getResp.Patient.BirthDate)
}
