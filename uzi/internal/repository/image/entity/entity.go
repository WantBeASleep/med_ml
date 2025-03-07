package entity

import (
	"uzi/internal/domain"

	"github.com/google/uuid"
)

type Image struct {
	Id    uuid.UUID `db:"id"`
	UziID uuid.UUID `db:"uzi_id"`
	Page  int       `db:"page"`
}

func (Image) FromDomain(d domain.Image) Image {
	return Image{
		Id:    d.Id,
		UziID: d.UziID,
		Page:  d.Page,
	}
}

func (Image) SliceFromDomain(slice []domain.Image) []Image {
	res := make([]Image, 0, len(slice))
	for _, v := range slice {
		res = append(res, Image{}.FromDomain(v))
	}
	return res
}

func (d Image) ToDomain() domain.Image {
	return domain.Image{
		Id:    d.Id,
		UziID: d.UziID,
		Page:  d.Page,
	}
}

func (Image) SliceToDomain(slice []Image) []domain.Image {
	res := make([]domain.Image, 0, len(slice))
	for _, v := range slice {
		res = append(res, v.ToDomain())
	}
	return res
}
