package echographic

import (
	"uzi/internal/repository/echographic/entity"
	repoEntity "uzi/internal/repository/entity"
)

func (q *repo) InsertEchographic(echographic entity.Echographic) error {
	query := q.QueryBuilder().
		Insert(table).
		Columns(
			columnID,
			columnContors,
			columnLeftLobeLength,
			columnLeftLobeWidth,
			columnLeftLobeThick,
			columnLeftLobeVolum,
			columnRightLobeLength,
			columnRightLobeWidth,
			columnRightLobeThick,
			columnRightLobeVolum,
			columnGlandVolum,
			columnIsthmus,
			columnStruct,
			columnEchogenicity,
			columnRegionalLymph,
			columnVascularization,
			columnLocation,
			columnAdditional,
			columnConclusion,
		).
		Values(
			echographic.Id,
			echographic.Contors,
			echographic.LeftLobeLength,
			echographic.LeftLobeWidth,
			echographic.LeftLobeThick,
			echographic.LeftLobeVolum,
			echographic.RightLobeLength,
			echographic.RightLobeWidth,
			echographic.RightLobeThick,
			echographic.RightLobeVolum,
			echographic.GlandVolum,
			echographic.Isthmus,
			echographic.Struct,
			echographic.Echogenicity,
			echographic.RegionalLymph,
			echographic.Vascularization,
			echographic.Location,
			echographic.Additional,
			echographic.Conclusion,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}

	return nil
}
