//go:build e2e

package doctor_test

import (
	pb "med/internal/generated/grpc/service"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func (suite *TestSuite) TestRegisterDoctor_Success() {
	id := uuid.New().String()
	fullname := gofakeit.Name()
	org := gofakeit.Company()
	job := gofakeit.JobTitle()
	description := gofakeit.MinecraftAnimal()

	_, err := suite.deps.Adapter.RegisterDoctor(
		suite.T().Context(),
		&pb.RegisterDoctorIn{
			Doctor: &pb.Doctor{
				Id:          id,
				Fullname:    fullname,
				Org:         org,
				Job:         job,
				Description: &description,
			},
		},
	)
	require.NoError(suite.T(), err)

	getResp, err := suite.deps.Adapter.GetDoctor(
		suite.T().Context(),
		&pb.GetDoctorIn{Id: id},
	)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), id, getResp.Doctor.Id)
	require.Equal(suite.T(), fullname, getResp.Doctor.Fullname)
	require.Equal(suite.T(), org, getResp.Doctor.Org)
	require.Equal(suite.T(), job, getResp.Doctor.Job)
	require.NotNil(suite.T(), getResp.Doctor.Description)
	require.Equal(suite.T(), description, *getResp.Doctor.Description)
}
