package patient

import (
	repoEntity "med/internal/repository/entity"
	entity "med/internal/repository/patient/entity"
)

func (r *repo) InsertPatient(patient entity.Patient) error {
	query := r.QueryBuilder().
		Insert(table).
		Columns(
			columnID,
			columnFullName,
			columnEmail,
			columnPolicy,
			columnActive,
			columnMalignancy,
			columnBirthDate,
			columnLastUziDate,
		).
		Values(
			patient.Id,
			patient.FullName,
			patient.Email,
			patient.Policy,
			patient.Active,
			patient.Malignancy,
			patient.BirthDate,
			patient.LastUziDate,
		)

	_, err := r.Runner().Execx(r.Context(), query)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}

	return nil
}
