package cytology_image

import (
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"cytology/internal/repository/cytology_image/entity"
	daoEntity "cytology/internal/repository/entity"
)

func (q *repo) GetCytologyImageByID(id uuid.UUID) (entity.CytologyImage, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnExternalID,
			columnDoctorID,
			columnPatientID,
			columnDiagnosticNumber,
			columnDiagnosticMarking,
			columnMaterialType,
			columnDiagnosDate,
			columnIsLast,
			columnCalcitonin,
			columnCalcitoninInFlush,
			columnThyroglobulin,
			columnDetails,
			columnPrevID,
			columnParentPrevID,
			columnCreateAt,
		).
		From(table).
		Where(sq.Eq{
			columnID: id,
		})

	var img entity.CytologyImage
	if err := q.Runner().Getx(q.Context(), &img, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.CytologyImage{}, daoEntity.ErrNotFound
		}
		return entity.CytologyImage{}, err
	}

	return img, nil
}

func (q *repo) GetCytologyImagesByExternalID(externalID uuid.UUID) ([]entity.CytologyImage, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnExternalID,
			columnDoctorID,
			columnPatientID,
			columnDiagnosticNumber,
			columnDiagnosticMarking,
			columnMaterialType,
			columnDiagnosDate,
			columnIsLast,
			columnCalcitonin,
			columnCalcitoninInFlush,
			columnThyroglobulin,
			columnDetails,
			columnPrevID,
			columnParentPrevID,
			columnCreateAt,
		).
		From(table).
		Where(sq.Eq{
			columnExternalID: externalID,
		})

	var images []entity.CytologyImage
	if err := q.Runner().Selectx(q.Context(), &images, query); err != nil {
		return nil, err
	}

	if len(images) == 0 {
		return nil, daoEntity.ErrNotFound
	}

	return images, nil
}

func (q *repo) GetCytologyImagesByDoctorIdAndPatientId(doctorID, patientID uuid.UUID) ([]entity.CytologyImage, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnExternalID,
			columnDoctorID,
			columnPatientID,
			columnDiagnosticNumber,
			columnDiagnosticMarking,
			columnMaterialType,
			columnDiagnosDate,
			columnIsLast,
			columnCalcitonin,
			columnCalcitoninInFlush,
			columnThyroglobulin,
			columnDetails,
			columnPrevID,
			columnParentPrevID,
			columnCreateAt,
		).
		From(table).
		Where(sq.Eq{
			columnDoctorID:  doctorID,
			columnPatientID: patientID,
		})

	var images []entity.CytologyImage
	if err := q.Runner().Selectx(q.Context(), &images, query); err != nil {
		return nil, err
	}

	if len(images) == 0 {
		return nil, daoEntity.ErrNotFound
	}

	return images, nil
}

func (q *repo) CheckExist(id uuid.UUID) (bool, error) {
	query := q.QueryBuilder().
		Select(columnID).
		Prefix("SELECT EXISTS (").
		From(table).
		Where(sq.Eq{
			columnID: id,
		}).
		Suffix(")")

	var exists bool
	if err := q.Runner().Getx(q.Context(), &exists, query); err != nil {
		return false, err
	}

	return exists, nil
}

func (q *repo) GetCytologyImagesByParentPrevID(parentPrevID uuid.UUID) ([]entity.CytologyImage, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnExternalID,
			columnDoctorID,
			columnPatientID,
			columnDiagnosticNumber,
			columnDiagnosticMarking,
			columnMaterialType,
			columnDiagnosDate,
			columnIsLast,
			columnCalcitonin,
			columnCalcitoninInFlush,
			columnThyroglobulin,
			columnDetails,
			columnPrevID,
			columnParentPrevID,
			columnCreateAt,
		).
		From(table).
		Where(sq.Eq{
			columnParentPrevID: parentPrevID,
		}).
		OrderBy(columnCreateAt + " ASC")

	var images []entity.CytologyImage
	if err := q.Runner().Selectx(q.Context(), &images, query); err != nil {
		return nil, err
	}

	return images, nil
}
