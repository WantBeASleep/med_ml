package repository

import (
	"context"

	daolib "github.com/WantBeASleep/med_ml_lib/dao"
	"github.com/jmoiron/sqlx"

	"med/internal/repository/card"
	"med/internal/repository/doctor"
	"med/internal/repository/patient"
)

type DAO interface {
	daolib.DAO
	NewDoctorQuery(ctx context.Context) doctor.Repository
	NewPatientQuery(ctx context.Context) patient.Repository
	NewCardQuery(ctx context.Context) card.Repository
}

type dao struct {
	daolib.DAO
}

func NewRepository(psql *sqlx.DB) DAO {
	return &dao{DAO: daolib.NewDao(psql)}
}

func (d *dao) NewDoctorQuery(ctx context.Context) doctor.Repository {
	doctorQuery := doctor.NewR()
	d.NewRepo(ctx, doctorQuery)

	return doctorQuery
}

func (d *dao) NewPatientQuery(ctx context.Context) patient.Repository {
	patientQuery := patient.NewR()
	d.NewRepo(ctx, patientQuery)

	return patientQuery
}

func (d *dao) NewCardQuery(ctx context.Context) card.Repository {
	cardQuery := card.NewR()
	d.NewRepo(ctx, cardQuery)

	return cardQuery
}
