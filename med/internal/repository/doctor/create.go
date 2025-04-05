package doctor

import (
	dentity "med/internal/repository/doctor/entity"
)

func (r *repo) InsertDoctor(doctor dentity.Doctor) error {
	query := r.QueryBuilder().
		Insert(table).
		Columns(
			columnID,
			columnFullname,
			columnOrg,
			columnJob,
			columnDescription,
		).
		Values(
			doctor.Id,
			doctor.FullName,
			doctor.Org,
			doctor.Job,
			doctor.Description,
		)

	_, err := r.Runner().Execx(r.Context(), query)
	if err != nil {
		return err
	}

	return nil
}
