package entity

import (
	"time"

	"cytology/internal/domain"
)

type Segmentation struct {
	Id                  int       `db:"id"`
	SegmentationGroupID int       `db:"segmentation_group_id"`
	CreateAt            time.Time `db:"create_at"`
	Points              []SegmentationPoint
}

type SegmentationPoint struct {
	Id             int       `db:"id"`
	SegmentationID int       `db:"segmentation_id"`
	X              int       `db:"x"`
	Y              int       `db:"y"`
	UID            int64     `db:"uid"`
	CreateAt       time.Time `db:"create_at"`
}

func (SegmentationPoint) FromDomain(d domain.SegmentationPoint) SegmentationPoint {
	return SegmentationPoint{
		Id:             d.Id,
		SegmentationID: d.SegmentationID,
		X:              d.X,
		Y:              d.Y,
		UID:            d.UID,
		CreateAt:       d.CreateAt,
	}
}

func (d SegmentationPoint) ToDomain() domain.SegmentationPoint {
	return domain.SegmentationPoint{
		Id:             d.Id,
		SegmentationID: d.SegmentationID,
		X:              d.X,
		Y:              d.Y,
		UID:            d.UID,
		CreateAt:       d.CreateAt,
	}
}

func (Segmentation) FromDomain(d domain.Segmentation) Segmentation {
	points := make([]SegmentationPoint, 0, len(d.Points))
	for _, p := range d.Points {
		points = append(points, SegmentationPoint{}.FromDomain(p))
	}

	return Segmentation{
		Id:                  d.Id,
		SegmentationGroupID: d.SegmentationGroupID,
		CreateAt:            d.CreateAt,
		Points:              points,
	}
}

func (d Segmentation) ToDomain() domain.Segmentation {
	points := make([]domain.SegmentationPoint, 0, len(d.Points))
	for _, p := range d.Points {
		points = append(points, p.ToDomain())
	}

	return domain.Segmentation{
		Id:                  d.Id,
		SegmentationGroupID: d.SegmentationGroupID,
		CreateAt:            d.CreateAt,
		Points:              points,
	}
}

func (Segmentation) SliceToDomain(segs []Segmentation) []domain.Segmentation {
	domainSegs := make([]domain.Segmentation, 0, len(segs))
	for _, v := range segs {
		domainSegs = append(domainSegs, v.ToDomain())
	}
	return domainSegs
}
