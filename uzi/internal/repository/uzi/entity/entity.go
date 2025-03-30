package entity

import (
	"database/sql"
	"time"

	"uzi/internal/domain"

	"github.com/WantBeASleep/med_ml_lib/gtc"
	"github.com/google/uuid"
)

type Uzi struct {
	Id          uuid.UUID      `db:"id"`
	Projection  string         `db:"projection"`
	Checked     bool           `db:"checked"`
	ExternalID  uuid.UUID      `db:"external_id"`
	Author      uuid.UUID      `db:"author"`
	DeviceID    int            `db:"device_id"`
	Status      string         `db:"status"`
	Description sql.NullString `db:"description"`
	CreateAt    time.Time      `db:"create_at"`
}

func (Uzi) FromDomain(d domain.Uzi) Uzi {
	return Uzi{
		Id:          d.Id,
		Projection:  d.Projection,
		Checked:     d.Checked,
		ExternalID:  d.ExternalID,
		Author:      d.Author,
		DeviceID:    d.DeviceID,
		Status:      d.Status.String(),
		Description: gtc.String.PointerToSql(d.Description),
		CreateAt:    d.CreateAt,
	}
}

func (d Uzi) ToDomain() domain.Uzi {
	// TODO: обработать ошибку
	// но нигде встретиться не должна
	status, _ := domain.UziStatus.Parse("", d.Status)

	return domain.Uzi{
		Id:          d.Id,
		Projection:  d.Projection,
		Checked:     d.Checked,
		ExternalID:  d.ExternalID,
		Author:      d.Author,
		DeviceID:    d.DeviceID,
		Status:      status,
		Description: gtc.String.SqlToPointer(d.Description),
		CreateAt:    d.CreateAt,
	}
}

func (Uzi) SliceToDomain(uzis []Uzi) []domain.Uzi {
	domainUzis := make([]domain.Uzi, 0, len(uzis))
	for _, v := range uzis {
		domainUzis = append(domainUzis, v.ToDomain())
	}
	return domainUzis
}
