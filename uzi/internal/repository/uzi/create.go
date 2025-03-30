package uzi

import (
	"uzi/internal/repository/uzi/entity"
)

func (q *repo) InsertUzi(uzi entity.Uzi) error {
	query := q.QueryBuilder().
		Insert(uziTable).
		Columns(
			columnID,
			columnProjection,
			columnChecked,
			columnExternalID,
			columnAuthor,
			columnDeviceID,
			columnStatus,
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
			uzi.CreateAt,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return nil
}
