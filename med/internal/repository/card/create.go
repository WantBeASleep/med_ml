package card

import (
	"med/internal/repository/card/entity"
	repoEntity "med/internal/repository/entity"
)

func (r *repo) InsertCard(card entity.Card) (int, error) {
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
		).
		Suffix("RETURNING id")

	var id int
	err := r.Runner().Getx(r.Context(), &id, query)
	if err != nil {
		return 0, repoEntity.WrapDBError(err)
	}

	return id, nil
}
