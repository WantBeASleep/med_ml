package card

import (
	"med/internal/repository/card/entity"
	repoEntity "med/internal/repository/entity"
)

func (r *repo) InsertCard(card entity.Card) error {
	query := r.QueryBuilder().
		Insert(table).
		Columns(
			columnDoctorID,
			columnPatientID,
			columnDiagnosis,
		).
		Values(
			card.DoctorID,
			card.PatientID,
			card.Diagnosis,
		)

	_, err := r.Runner().Execx(r.Context(), query)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}

	return nil
}
