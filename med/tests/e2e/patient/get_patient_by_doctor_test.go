//go:build e2e

package patient_test

import (
	"time"

	pb "med/internal/generated/grpc/service"
	"med/tests/e2e/flow"

	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestGetPatientByDoctorID_Success() {
	data, err := flow.New(suite.deps,
		flow.CreatePatient,
		flow.CreateDoctor,
		flow.CreateCard,
	).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	getResp, err := suite.deps.Adapter.GetPatientsByDoctorID(
		suite.T().Context(),
		&pb.GetPatientsByDoctorIDIn{Id: data.Doctor.Id.String()},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 1, len(getResp.Patients))
	require.Equal(suite.T(), data.Patient.Id.String(), getResp.Patients[0].Id)
	require.Equal(suite.T(), data.Patient.FullName, getResp.Patients[0].Fullname)
	require.Equal(suite.T(), data.Patient.Email, getResp.Patients[0].Email)
	require.Equal(suite.T(), data.Patient.Policy, getResp.Patients[0].Policy)
	require.Equal(suite.T(), data.Patient.Active, getResp.Patients[0].Active)
	require.Equal(suite.T(), data.Patient.Malignancy, getResp.Patients[0].Malignancy)
	require.Equal(suite.T(), data.Patient.BirthDate.Format(time.RFC3339), getResp.Patients[0].BirthDate)
}
