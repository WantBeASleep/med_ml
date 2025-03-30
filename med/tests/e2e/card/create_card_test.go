//go:build e2e

package card_test

import (
	"med/tests/e2e/flow"

	pb "med/internal/generated/grpc/service"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestCreateCard_Success() {
	data, err := flow.New(suite.deps,
		flow.CreatePatient,
		flow.CreateDoctor,
	).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	diagnosis := gofakeit.MinecraftAnimal()

	_, err = suite.deps.Adapter.CreateCard(
		suite.T().Context(),
		&pb.CreateCardIn{
			Card: &pb.Card{
				PatientId: data.Patient.Id.String(),
				DoctorId:  data.Doctor.Id.String(),
				Diagnosis: &diagnosis,
			},
		},
	)
	require.NoError(suite.T(), err)

	getCardResp, err := suite.deps.Adapter.GetCard(
		suite.T().Context(),
		&pb.GetCardIn{
			DoctorId:  data.Doctor.Id.String(),
			PatientId: data.Patient.Id.String(),
		},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), data.Patient.Id.String(), getCardResp.Card.PatientId)
	require.Equal(suite.T(), data.Doctor.Id.String(), getCardResp.Card.DoctorId)
	require.Equal(suite.T(), diagnosis, *getCardResp.Card.Diagnosis)
}
