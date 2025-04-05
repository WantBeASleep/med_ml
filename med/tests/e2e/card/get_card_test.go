//go:build e2e

package card_test

import (
	pb "med/internal/generated/grpc/service"
	"med/tests/e2e/flow"

	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestGetCard_Success() {
	data, err := flow.New(suite.deps,
		flow.CreatePatient,
		flow.CreateDoctor,
		flow.CreateCard,
	).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	getResp, err := suite.deps.Adapter.GetCard(
		suite.T().Context(),
		&pb.GetCardIn{
			DoctorId:  data.Doctor.Id.String(),
			PatientId: data.Patient.Id.String(),
		},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), data.Patient.Id.String(), getResp.Card.PatientId)
	require.Equal(suite.T(), data.Doctor.Id.String(), getResp.Card.DoctorId)
	require.Equal(suite.T(), *data.Card.Diagnosis, *getResp.Card.Diagnosis)
}
