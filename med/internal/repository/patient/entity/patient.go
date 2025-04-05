package entity

import (
	"database/sql"
	"time"

	gtclib "github.com/WantBeASleep/med_ml_lib/gtc"

	"med/internal/domain"

	"github.com/google/uuid"
)

type Patient struct {
	Id          uuid.UUID    `db:"id"`
	FullName    string       `db:"fullname"`
	Email       string       `db:"email"`
	Policy      string       `db:"policy"`
	Active      bool         `db:"active"`
	Malignancy  bool         `db:"malignancy"`
	BirthDate   time.Time    `db:"birth_date"`
	LastUziDate sql.NullTime `db:"last_uzi_date"`
}

func (Patient) FromDomain(p domain.Patient) Patient {
	return Patient{
		Id:          p.Id,
		FullName:    p.FullName,
		Email:       p.Email,
		Policy:      p.Policy,
		Active:      p.Active,
		Malignancy:  p.Malignancy,
		BirthDate:   p.BirthDate,
		LastUziDate: gtclib.Time.PointerToSql(p.LastUziDate),
	}
}

func (p Patient) ToDomain() domain.Patient {
	return domain.Patient{
		Id:          p.Id,
		FullName:    p.FullName,
		Email:       p.Email,
		Policy:      p.Policy,
		Active:      p.Active,
		Malignancy:  p.Malignancy,
		BirthDate:   p.BirthDate,
		LastUziDate: gtclib.Time.SqlToPointer(p.LastUziDate),
	}
}

func (p Patient) SliceToDomain(patients []Patient) []domain.Patient {
	res := make([]domain.Patient, 0, len(patients))
	for _, v := range patients {
		res = append(res, v.ToDomain())
	}
	return res
}
