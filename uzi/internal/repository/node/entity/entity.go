package entity

import (
	"database/sql"

	"github.com/WantBeASleep/med_ml_lib/gtc"
	"github.com/google/uuid"

	"uzi/internal/domain"
)

type Node struct {
	Id         uuid.UUID      `db:"id"`
	Ai         bool           `db:"ai"`
	UziID      uuid.UUID      `db:"uzi_id"`
	Validation sql.NullString `db:"validation"`
	Tirads23   float64        `db:"tirads_23"`
	Tirads4    float64        `db:"tirads_4"`
	Tirads5    float64        `db:"tirads_5"`
}

func (Node) FromDomain(d domain.Node) Node {
	return Node{
		Id:         d.Id,
		Ai:         d.Ai,
		UziID:      d.UziID,
		Validation: gtc.String.PointerToSql((*string)(d.Validation)),
		Tirads23:   d.Tirads23,
		Tirads4:    d.Tirads4,
		Tirads5:    d.Tirads5,
	}
}

func (Node) SliceFromDomain(slice []domain.Node) []Node {
	res := make([]Node, 0, len(slice))
	for _, v := range slice {
		res = append(res, Node{}.FromDomain(v))
	}
	return res
}

func (d Node) ToDomain() domain.Node {
	node := domain.Node{
		Id:       d.Id,
		Ai:       d.Ai,
		UziID:    d.UziID,
		Tirads23: d.Tirads23,
		Tirads4:  d.Tirads4,
		Tirads5:  d.Tirads5,
	}

	if d.Validation.Valid {
		validation, _ := domain.NodeValidation.Parse("", d.Validation.String)
		node.Validation = &validation
	}

	return node
}

func (Node) SliceToDomain(slice []Node) []domain.Node {
	res := make([]domain.Node, 0, len(slice))
	for _, v := range slice {
		res = append(res, v.ToDomain())
	}
	return res
}
