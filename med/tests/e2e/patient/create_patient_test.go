//go:build e2e

package patient_test

import (
	"time"

	pb "med/internal/generated/grpc/service"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestCreatePatient_Success() {
	fullname := gofakeit.Name()
	email := gofakeit.Email()
	policy := gofakeit.MinecraftAnimal()
	active := gofakeit.Bool()
	malignancy := gofakeit.Bool()
	birthDate := gofakeit.Date()

	resp, err := suite.deps.Adapter.CreatePatient(
		suite.T().Context(),
		&pb.CreatePatientIn{
			Fullname:   fullname,
			Email:      email,
			Policy:     policy,
			Active:     active,
			Malignancy: malignancy,
			BirthDate:  birthDate.Format(time.RFC3339),
		},
	)
	require.NoError(suite.T(), err)

	getResp, err := suite.deps.Adapter.GetPatient(
		suite.T().Context(),
		&pb.GetPatientIn{Id: resp.Id},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), resp.Id, getResp.Patient.Id)
	require.Equal(suite.T(), fullname, getResp.Patient.Fullname)
	require.Equal(suite.T(), email, getResp.Patient.Email)
	require.Equal(suite.T(), policy, getResp.Patient.Policy)
	require.Equal(suite.T(), active, getResp.Patient.Active)
	require.Equal(suite.T(), malignancy, getResp.Patient.Malignancy)
	require.Equal(suite.T(), birthDate.Format(time.RFC3339), getResp.Patient.BirthDate)
}
