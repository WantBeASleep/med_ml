package device

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
		return 0, err
	}

	return id, nil
}
