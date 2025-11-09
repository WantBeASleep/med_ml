package echographic

import (
	"uzi/internal/repository/echographic/entity"
	repoEntity "uzi/internal/repository/entity"

	sq "github.com/Masterminds/squirrel"
)

func (q *repo) UpdateEchographic(echographic entity.Echographic) error {
	query := q.QueryBuilder().
		Update(table).
		SetMap(sq.Eq{
			columnContors:         echographic.Contors,
			columnLeftLobeLength:  echographic.LeftLobeLength,
			columnLeftLobeWidth:   echographic.LeftLobeWidth,
			columnLeftLobeThick:   echographic.LeftLobeThick,
			columnLeftLobeVolum:   echographic.LeftLobeVolum,
			columnRightLobeLength: echographic.RightLobeLength,
			columnRightLobeWidth:  echographic.RightLobeWidth,
			columnRightLobeThick:  echographic.RightLobeThick,
			columnRightLobeVolum:  echographic.RightLobeVolum,
			columnGlandVolum:      echographic.GlandVolum,
			columnIsthmus:         echographic.Isthmus,
			columnStruct:          echographic.Struct,
			columnEchogenicity:    echographic.Echogenicity,
			columnRegionalLymph:   echographic.RegionalLymph,
			columnVascularization: echographic.Vascularization,
			columnLocation:        echographic.Location,
			columnAdditional:      echographic.Additional,
			columnConclusion:      echographic.Conclusion,
		}).
		Where(sq.Eq{
			columnID: echographic.Id,
		})

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}

	return nil
}
