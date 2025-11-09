package node

import (
	repoEntity "uzi/internal/repository/entity"
	"uzi/internal/repository/node/entity"

	sq "github.com/Masterminds/squirrel"
)

func (q *repo) UpdateNode(node entity.Node) error {
	query := q.QueryBuilder().
		Update(table).
		SetMap(sq.Eq{
			columnValidation: node.Validation,
			columnTirads23:   node.Tirads23,
			columnTirads4:    node.Tirads4,
			columnTirads5:    node.Tirads5,
		}).
		Where(sq.Eq{
			columnID: node.Id,
		})

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}

	return nil
}
