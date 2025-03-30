package card

import (
	"med/internal/repository/card/entity"

	sq "github.com/Masterminds/squirrel"
)

func (r *repo) UpdateCard(card entity.Card) error {
	query := r.QueryBuilder().
		Update(table).
		SetMap(sq.Eq{
			columnDiagnosis: card.Diagnosis,
		}).
		Where(sq.Eq{
			columnDoctorID:  card.DoctorID,
			columnPatientID: card.PatientID,
		})

	_, err := r.Runner().Execx(r.Context(), query)
	return err
}
