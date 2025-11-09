package segment

import (
	repoEntity "uzi/internal/repository/entity"
	"uzi/internal/repository/segment/entity"
)

func (q *repo) InsertSegments(segments ...entity.Segment) error {
	query := q.QueryBuilder().
		Insert(table).
		Columns(
			columnID,
			columnNodeID,
			columnImageID,
			columnContor,
			columnAi,
			columnTirads23,
			columnTirads4,
			columnTirads5,
		)

	for _, v := range segments {
		query = query.Values(
			v.Id,
			v.NodeID,
			v.ImageID,
			v.Contor,
			v.Ai,
			v.Tirads23,
			v.Tirads4,
			v.Tirads5,
		)
	}
	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}

	return nil
}
