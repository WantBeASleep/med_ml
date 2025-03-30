package server

import (
	"med/internal/generated/grpc/service"
	"med/internal/server/card"
	"med/internal/server/doctor"
	"med/internal/server/patient"
	"med/internal/services"
)

type Server struct {
	patient.PatientHandler
	doctor.DoctorHandler
	card.CardHandler

	service.UnsafeMedSrvServer
}

func New(
	service *services.Service,
) *Server {
	patientHandler := patient.New(service.PatientService)
	doctorHandler := doctor.New(service.DoctorService)
	cardHandler := card.New(service.CardService)

	return &Server{
		PatientHandler: patientHandler,
		DoctorHandler:  doctorHandler,
		CardHandler:    cardHandler,
	}
}
