package entity

import (
	"database/sql"
	"encoding/json"
	"time"

	"cytology/internal/domain"

	"github.com/google/uuid"
)

type SegmentationGroup struct {
	Id         int            `db:"id"`
	CytologyID uuid.UUID      `db:"cytology_id"`
	SegType    string         `db:"seg_type"`
	GroupType  string         `db:"group_type"`
	IsAI       bool           `db:"is_ai"`
	Details    sql.NullString `db:"details"`
	CreateAt   time.Time      `db:"create_at"`
}

func (SegmentationGroup) FromDomain(d domain.SegmentationGroup) SegmentationGroup {
	var details sql.NullString
	if len(d.Details) > 0 && string(d.Details) != "null" {
		details = sql.NullString{String: string(d.Details), Valid: true}
	}

	return SegmentationGroup{
		Id:         d.Id,
		CytologyID: d.CytologyID,
		SegType:    d.SegType.String(),
		GroupType:  d.GroupType.String(),
		IsAI:       d.IsAI,
		Details:    details,
		CreateAt:   d.CreateAt,
	}
}

func (d SegmentationGroup) ToDomain() domain.SegmentationGroup {
	var segType domain.SegType
	var groupType domain.GroupType
	// Простой парсинг, можно улучшить
	segType = domain.SegType(d.SegType)
	groupType = domain.GroupType(d.GroupType)

	var details json.RawMessage
	if d.Details.Valid {
		details = json.RawMessage(d.Details.String)
	}

	return domain.SegmentationGroup{
		Id:         d.Id,
		CytologyID: d.CytologyID,
		SegType:    segType,
		GroupType:  groupType,
		IsAI:       d.IsAI,
		Details:    details,
		CreateAt:   d.CreateAt,
	}
}

func (SegmentationGroup) SliceToDomain(groups []SegmentationGroup) []domain.SegmentationGroup {
	domainGroups := make([]domain.SegmentationGroup, 0, len(groups))
	for _, v := range groups {
		domainGroups = append(domainGroups, v.ToDomain())
	}
	return domainGroups
}
