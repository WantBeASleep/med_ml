package entity

import (
	"database/sql"

	gtclib "github.com/WantBeASleep/med_ml_lib/gtc"
	"github.com/google/uuid"

	"med/internal/domain"
)

type Doctor struct {
	Id          uuid.UUID      `db:"id"`
	FullName    string         `db:"fullname"`
	Org         string         `db:"org"`
	Job         string         `db:"job"`
	Description sql.NullString `db:"description"`
}

func (Doctor) FromDomain(d domain.Doctor) Doctor {
	return Doctor{
		Id:          d.Id,
		FullName:    d.FullName,
		Org:         d.Org,
		Job:         d.Job,
		Description: gtclib.String.PointerToSql(d.Description),
	}
}

func (d Doctor) ToDomain() domain.Doctor {
	return domain.Doctor{
		Id:          d.Id,
		FullName:    d.FullName,
		Org:         d.Org,
		Job:         d.Job,
		Description: gtclib.String.SqlToPointer(d.Description),
	}
}
