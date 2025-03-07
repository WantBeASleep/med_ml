package entity

import (
	"time"

	"uzi/internal/domain"

	"github.com/google/uuid"
)

type Uzi struct {
	Id         uuid.UUID `db:"id"`
	Projection string    `db:"projection"`
	Checked    bool      `db:"checked"`
	ExternalID uuid.UUID `db:"external_id"`
	DeviceID   int       `db:"device_id"`
	Status     string    `db:"status"`
	CreateAt   time.Time `db:"create_at"`
}

func (Uzi) FromDomain(d domain.Uzi) Uzi {
	return Uzi{
		Id:         d.Id,
		Projection: d.Projection,
		Checked:    d.Checked,
		ExternalID: d.ExternalID,
		DeviceID:   d.DeviceID,
		Status:     d.Status.String(),
		CreateAt:   d.CreateAt,
	}
}

func (d Uzi) ToDomain() domain.Uzi {
	// TODO: обработать ошибку
	// но нигде встретиться не должна
	status, _ := domain.UziStatus.Parse("", d.Status)

	return domain.Uzi{
		Id:         d.Id,
		Projection: d.Projection,
		Checked:    d.Checked,
		ExternalID: d.ExternalID,
		DeviceID:   d.DeviceID,
		Status:     status,
		CreateAt:   d.CreateAt,
	}
}

func (Uzi) SliceToDomain(uzis []Uzi) []domain.Uzi {
	domainUzis := make([]domain.Uzi, 0, len(uzis))
	for _, v := range uzis {
		domainUzis = append(domainUzis, v.ToDomain())
	}
	return domainUzis
}
