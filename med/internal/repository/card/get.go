package card

import (
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"med/internal/repository/card/entity"
	daoEntity "med/internal/repository/entity"
)

func (r *repo) GetCardByPK(doctorID uuid.UUID, patientID uuid.UUID) (entity.Card, error) {
	query := r.QueryBuilder().
		Select(
			columnID,
			columnDoctorID,
			columnPatientID,
			columnDiagnosis,
		).
		From(table).
		Where(sq.Eq{
			columnDoctorID:  doctorID,
			columnPatientID: patientID,
		})

	var card entity.Card
	if err := r.Runner().Getx(r.Context(), &card, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Card{}, daoEntity.ErrNotFound
		}
		return entity.Card{}, err
	}

	return card, nil
}
