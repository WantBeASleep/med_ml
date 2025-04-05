//go:build e2e

package card_test

import (
	pb "med/internal/generated/grpc/service"
	"med/tests/e2e/flow"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestUpdateCard_Success() {
	data, err := flow.New(suite.deps,
		flow.CreatePatient,
		flow.CreateDoctor,
		flow.CreateCard,
	).Do(suite.T().Context())
	require.NoError(suite.T(), err)

	diagnosis := gofakeit.MinecraftAnimal()

	updateResp, err := suite.deps.Adapter.UpdateCard(
		suite.T().Context(),
		&pb.UpdateCardIn{
			Card: &pb.Card{
				PatientId: data.Patient.Id.String(),
				DoctorId:  data.Doctor.Id.String(),
				Diagnosis: &diagnosis,
			},
		},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), data.Patient.Id.String(), updateResp.Card.PatientId)
	require.Equal(suite.T(), data.Doctor.Id.String(), updateResp.Card.DoctorId)
	require.Equal(suite.T(), diagnosis, *updateResp.Card.Diagnosis)
}
