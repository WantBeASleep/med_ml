package cytology_image

import (
	"cytology/internal/repository/cytology_image/entity"
	repoEntity "cytology/internal/repository/entity"
)

func (q *repo) InsertCytologyImage(img entity.CytologyImage) error {
	query := q.QueryBuilder().
		Insert(table).
		Columns(
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
		Values(
			img.Id,
			img.ExternalID,
			img.DoctorID,
			img.PatientID,
			img.DiagnosticNumber,
			img.DiagnosticMarking,
			img.MaterialType,
			img.DiagnosDate,
			img.IsLast,
			img.Calcitonin,
			img.CalcitoninInFlush,
			img.Thyroglobulin,
			img.Details,
			img.PrevID,
			img.ParentPrevID,
			img.CreateAt,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return repoEntity.WrapDBError(err)
	}

	return nil
}
