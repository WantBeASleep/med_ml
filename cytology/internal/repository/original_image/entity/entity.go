package entity

import (
	"database/sql"
	"time"

	"cytology/internal/domain"

	"github.com/WantBeASleep/med_ml_lib/gtc"
	"github.com/google/uuid"
)

type OriginalImage struct {
	Id         uuid.UUID      `db:"id"`
	CytologyID uuid.UUID      `db:"cytology_id"`
	ImagePath  string         `db:"image_path"`
	CreateDate time.Time      `db:"create_date"`
	DelayTime  sql.NullFloat64 `db:"delay_time"`
	ViewedFlag bool           `db:"viewed_flag"`
}

func (OriginalImage) FromDomain(d domain.OriginalImage) OriginalImage {
	return OriginalImage{
		Id:         d.Id,
		CytologyID: d.CytologyID,
		ImagePath:  d.ImagePath,
		CreateDate: d.CreateDate,
		DelayTime:  gtc.Float64.PointerToSql(d.DelayTime),
		ViewedFlag: d.ViewedFlag,
	}
}

func (d OriginalImage) ToDomain() domain.OriginalImage {
	return domain.OriginalImage{
		Id:         d.Id,
		CytologyID: d.CytologyID,
		ImagePath:  d.ImagePath,
		CreateDate: d.CreateDate,
		DelayTime:  gtc.Float64.SqlToPointer(d.DelayTime),
		ViewedFlag: d.ViewedFlag,
	}
}

func (OriginalImage) SliceToDomain(images []OriginalImage) []domain.OriginalImage {
	domainImages := make([]domain.OriginalImage, 0, len(images))
	for _, v := range images {
		domainImages = append(domainImages, v.ToDomain())
	}
	return domainImages
}
