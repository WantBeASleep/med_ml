package patient

import (
	"database/sql"
	"errors"
	"fmt"

	entity "med/internal/repository/entity"
	pentity "med/internal/repository/patient/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (r *repo) GetPatientByID(id uuid.UUID) (pentity.Patient, error) {
	query := r.QueryBuilder().
		Select(
			columnID,
			columnFullName,
			columnEmail,
			columnPolicy,
			columnActive,
			columnMalignancy,
			columnBirthDate,
			columnLastUziDate,
		).
		From(table).
		Where(sq.Eq{
			columnID: id,
		})

	var patient pentity.Patient
	if err := r.Runner().Getx(r.Context(), &patient, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pentity.Patient{}, entity.ErrNotFound
		}
		return pentity.Patient{}, err
	}

	return patient, nil
}

func (r *repo) GetPatientsByDoctorID(id uuid.UUID) ([]pentity.Patient, error) {
	query := r.QueryBuilder().
		Select(
			fmt.Sprintf("%s.%s", table, columnID),
			fmt.Sprintf("%s.%s", table, columnFullName),
			fmt.Sprintf("%s.%s", table, columnEmail),
			fmt.Sprintf("%s.%s", table, columnPolicy),
			fmt.Sprintf("%s.%s", table, columnActive),
			fmt.Sprintf("%s.%s", table, columnMalignancy),
			fmt.Sprintf("%s.%s", table, columnBirthDate),
			fmt.Sprintf("%s.%s", table, columnLastUziDate),
		).
		From(table).
		InnerJoin("card ON card.patient_id = patient.id").
		Where(sq.Eq{
			"card.doctor_id": id,
		})

	var patient []pentity.Patient
	if err := r.Runner().Selectx(r.Context(), &patient, query); err != nil {
		return nil, err
	}

	if len(patient) == 0 {
		return nil, entity.ErrNotFound
	}

	return patient, nil
}
