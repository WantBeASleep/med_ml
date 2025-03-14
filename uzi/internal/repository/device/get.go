package device

import (
	"uzi/internal/repository/device/entity"
)

func (q *repo) GetDeviceList() ([]entity.Device, error) {
	query := q.QueryBuilder().
		Select(
			columnId,
			columnName,
		).
		From(table)

	var devices []entity.Device
	if err := q.Runner().Selectx(q.Context(), &devices, query); err != nil {
		return nil, err
	}

	return devices, nil
}
