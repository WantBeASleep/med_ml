package node

import (
	"uzi/internal/repository/node/entity"
)

func (q *repo) InsertNodes(nodes ...entity.Node) error {
	query := q.QueryBuilder().
		Insert(table).
		Columns(
			columnID,
			columnAI,
			columnUziID,
			columnValidation,
			columnTirads23,
			columnTirads4,
			columnTirads5,
			columnDescription,
		)

	for _, v := range nodes {
		query = query.Values(
			v.Id,
			v.Ai,
			v.UziID,
			v.Validation,
			v.Tirads23,
			v.Tirads4,
			v.Tirads5,
			v.Description,
		)
	}
	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return nil
}
