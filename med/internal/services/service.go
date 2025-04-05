package services

import (
	"med/internal/repository"
	"med/internal/services/card"
	"med/internal/services/doctor"
	"med/internal/services/patient"
)

type Service struct {
	CardService    card.Service
	DoctorService  doctor.Service
	PatientService patient.Service
}

func NewService(
	dao repository.DAO,
) *Service {
	cardSrv := card.New(dao)
	doctorSrv := doctor.New(dao)
	patientSrv := patient.New(dao)

	return &Service{
		CardService:    cardSrv,
		DoctorService:  doctorSrv,
		PatientService: patientSrv,
	}
}
