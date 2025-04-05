package med

import (
	"composition-api/internal/server/med/card"
	"composition-api/internal/server/med/doctor"
	"composition-api/internal/server/med/patient"
	services "composition-api/internal/services"
)

type MedRoute interface {
	patient.PatientHandler
	doctor.DoctorHandler
	card.CardHandler
}

type medRoute struct {
	patient.PatientHandler
	doctor.DoctorHandler
	card.CardHandler
}

func NewMedRoute(services *services.Services) MedRoute {
	patientHandler := patient.NewHandler(services)
	doctorHandler := doctor.NewHandler(services)
	cardHandler := card.NewHandler(services)

	return &medRoute{
		patientHandler,
		doctorHandler,
		cardHandler,
	}
}
