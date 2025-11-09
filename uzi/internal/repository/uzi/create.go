package uzi

import (
	repoEntity "uzi/internal/repository/entity"
	"uzi/internal/repository/uzi/entity"
)

func (q *repo) InsertUzi(uzi entity.Uzi) error {
	query := q.QueryBuilder().
		Insert(table).
		Columns(
			columnID,
			columnProjection,
			columnChecked,
			columnExternalID,
			columnAuthor,
			columnDeviceID,
			columnStatus,
			columnDescription,
			columnCreateAt,
		).
		Values(
			uzi.Id,
			uzi.Projection,
			uzi.Checked,
			uzi.ExternalID,
			uzi.Author,
			uzi.DeviceID,
			uzi.Status,
			uzi.Description,
			uzi.CreateAt,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}

	return nil
}
