package patient

import (
	entity "med/internal/repository/patient/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r *repo) UpdatePatient(patient entity.Patient) error {
	query := r.QueryBuilder().
		Update(table).
		SetMap(sq.Eq{
			columnActive:      patient.Active,
			columnMalignancy:  patient.Malignancy,
			columnLastUziDate: patient.LastUziDate,
		}).
		Where(sq.Eq{
			columnID: patient.Id,
		})

	_, err := r.Runner().Execx(r.Context(), query)
	return err
}
