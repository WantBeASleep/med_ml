package device

import (
	repoEntity "uzi/internal/repository/entity"
)

func (q *repo) CreateDevice(name string) (int, error) {
	query := q.QueryBuilder().
		Insert(table).
		Columns(
			columnName,
		).
		Values(
			name,
		).
		Suffix("RETURNING id")

	var id int
	if err := q.Runner().Getx(q.Context(), &id, query); err != nil {
		return 0, repoEntity.WrapDBError(err)
	}

	return id, nil
}
