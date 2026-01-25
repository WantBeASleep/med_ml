package mappers

import (
	"github.com/google/uuid"

	domain "composition-api/internal/domain/cytology"
	med_domain "composition-api/internal/domain/med"
	api "composition-api/internal/generated/http/api"
	cytologySrv "composition-api/internal/services/cytology"
)

type CytologyImage struct{}

// Удалены неиспользуемые методы Domain, SliceDomain, CreateArg - заменены на новые методы для работы с обновленными типами API

func (CytologyImage) CreateArgFromCytologyCreateCreateReq(req *api.CytologyCreateCreateReq) cytologySrv.CreateCytologyImageArg {
	arg := cytologySrv.CreateCytologyImageArg{
		DiagnosticNumber: req.DiagnosticNumber,
		// ExternalID, DoctorID, PatientID должны быть получены из контекста или другого источника
		// Пока используем пустые UUID, но это нужно будет исправить
		ExternalID: uuid.Nil,
		DoctorID:   uuid.Nil,
		PatientID:  uuid.Nil,
	}

	// Обработка файла изображения
	if req.Image.Set {
		arg.File = &req.Image.Value
		// Получаем Content-Type из заголовка файла
		contentType := req.Image.Value.Header.Get("Content-Type")
		if contentType == "" {
			contentType = "application/octet-stream"
		}
		arg.ContentType = contentType
	}

	if req.DiagnosticMarking.Set {
		marking := domain.DiagnosticMarking(req.DiagnosticMarking.Value)
		arg.DiagnosticMarking = &marking
	}

	if req.MaterialType.Set {
		materialType := domain.MaterialType(req.MaterialType.Value)
		arg.MaterialType = &materialType
	}

	if req.Calcitonin.Set {
		calcitonin := int(req.Calcitonin.Value)
		arg.Calcitonin = &calcitonin
	}

	if req.CalcitoninInFlush.Set {
		calcitoninInFlush := int(req.CalcitoninInFlush.Value)
		arg.CalcitoninInFlush = &calcitoninInFlush
	}

	if req.Thyroglobulin.Set {
		thyroglobulin := int(req.Thyroglobulin.Value)
		arg.Thyroglobulin = &thyroglobulin
	}

	if req.Details.Set {
		// Details handling - строка JSON
		detailsStr := req.Details.Value
		arg.Details = &detailsStr
	}

	// Prev и ParentPrev теперь UUID (строки) в API, соответствуют UUID в БД

	return arg
}

// Удален неиспользуемый метод UpdateArg - заменен на UpdateArgFromCytologyUpdateUpdateReq и UpdateArgFromCytologyUpdatePartialUpdateReq

func (CytologyImage) ToCytologyReadOKInfo(img domain.CytologyImage, patient med_domain.Patient, patientCard med_domain.Card) api.CytologyReadOKInfo {
	// Маппим данные о пациенте
	apiPatient := api.Patient{
		ID:         patient.Id,
		Fullname:   patient.FullName,
		Email:      patient.Email,
		Policy:     patient.Policy,
		Active:     patient.Active,
		Malignancy: patient.Malignancy,
		BirthDate:  patient.BirthDate,
	}
	if patient.LastUziDate != nil {
		apiPatient.LastUziDate = api.OptDate{
			Value: *patient.LastUziDate,
			Set:   true,
		}
	}

	// Маппим данные о карточке пациента
	apiPatientCard := api.PatientCard{
		Patient: api.OptInt{
			// TODO: Преобразовать UUID в int (нужен lookup или другой способ)
			Set: false,
		},
		MedWorker: api.OptInt{
			// TODO: Преобразовать UUID в int (нужен lookup или другой способ)
			Set: false,
		},
		Diagnosis: api.OptString{
			Value: "",
			Set:   false,
		},
	}
	if patientCard.Diagnosis != nil {
		apiPatientCard.Diagnosis = api.OptString{
			Value: *patientCard.Diagnosis,
			Set:   true,
		}
	}

	imageGroup := api.CytologyReadOKInfoImageGroup{
		DiagnosticNumber: int(img.DiagnosticNumber),
	}

	if img.DiagnosticMarking != nil {
		imageGroup.DiagnosticMarking = api.OptCytologyReadOKInfoImageGroupDiagnosticMarking{
			Value: api.CytologyReadOKInfoImageGroupDiagnosticMarking(*img.DiagnosticMarking),
			Set:   true,
		}
	}

	if img.MaterialType != nil {
		imageGroup.MaterialType = api.OptCytologyReadOKInfoImageGroupMaterialType{
			Value: api.CytologyReadOKInfoImageGroupMaterialType(*img.MaterialType),
			Set:   true,
		}
	}

	if img.Calcitonin != nil {
		imageGroup.Calcitonin = api.OptInt{
			Value: int(*img.Calcitonin),
			Set:   true,
		}
	}

	if img.CalcitoninInFlush != nil {
		imageGroup.CalcitoninInFlush = api.OptInt{
			Value: int(*img.CalcitoninInFlush),
			Set:   true,
		}
	}

	if img.Thyroglobulin != nil {
		imageGroup.Thyroglobulin = api.OptInt{
			Value: int(*img.Thyroglobulin),
			Set:   true,
		}
	}

	imageGroup.IsLast = api.OptBool{
		Value: img.IsLast,
		Set:   true,
	}

	imageGroup.DiagnosDate = api.OptDateTime{
		Value: img.DiagnosDate,
		Set:   true,
	}

	// Маппинг prev и parent_prev
	if img.PrevID != nil {
		imageGroup.Prev = api.OptUUID{
			Value: *img.PrevID,
			Set:   true,
		}
	}
	if img.ParentPrevID != nil {
		imageGroup.ParentPrev = api.OptUUID{
			Value: *img.ParentPrevID,
			Set:   true,
		}
	}

	return api.CytologyReadOKInfo{
		Patient:     apiPatient,
		PatientCard: apiPatientCard,
		ImageGroup:  imageGroup,
	}
}

func (CytologyImage) ToCytologyImageModelList(imgs []domain.CytologyImage) []api.CytologyHistoryReadOKResultsItem {
	result := make([]api.CytologyHistoryReadOKResultsItem, 0, len(imgs))
	for _, img := range imgs {
		item := api.CytologyHistoryReadOKResultsItem{
			DiagnosticNumber: int(img.DiagnosticNumber),
			IsLast: api.OptBool{
				Value: img.IsLast,
				Set:   true,
			},
			DiagnosDate: api.OptDateTime{
				Value: img.DiagnosDate,
				Set:   true,
			},
		}

		if img.DiagnosticMarking != nil {
			item.DiagnosticMarking = api.OptCytologyHistoryReadOKResultsItemDiagnosticMarking{
				Value: api.CytologyHistoryReadOKResultsItemDiagnosticMarking(*img.DiagnosticMarking),
				Set:   true,
			}
		}

		if img.MaterialType != nil {
			item.MaterialType = api.OptCytologyHistoryReadOKResultsItemMaterialType{
				Value: api.CytologyHistoryReadOKResultsItemMaterialType(*img.MaterialType),
				Set:   true,
			}
		}

		if img.Calcitonin != nil {
			item.Calcitonin = api.OptInt{
				Value: int(*img.Calcitonin),
				Set:   true,
			}
		}

		if img.CalcitoninInFlush != nil {
			item.CalcitoninInFlush = api.OptInt{
				Value: int(*img.CalcitoninInFlush),
				Set:   true,
			}
		}

		if img.Thyroglobulin != nil {
			item.Thyroglobulin = api.OptInt{
				Value: int(*img.Thyroglobulin),
				Set:   true,
			}
		}

		if img.Details != nil {
			item.Details = &api.CytologyHistoryReadOKResultsItemDetails{}
		}

		// Маппинг prev и parent_prev
		if img.PrevID != nil {
			item.Prev = api.OptUUID{
				Value: *img.PrevID,
				Set:   true,
			}
		}
		if img.ParentPrevID != nil {
			item.ParentPrev = api.OptUUID{
				Value: *img.ParentPrevID,
				Set:   true,
			}
		}

		result = append(result, item)
	}
	return result
}

func (CytologyImage) UpdateArgFromCytologyUpdateUpdateReq(id uuid.UUID, req *api.CytologyUpdateUpdateReq) cytologySrv.UpdateCytologyImageArg {
	arg := cytologySrv.UpdateCytologyImageArg{
		Id: id,
	}

	// Извлекаем данные из req.Details или из верхнего уровня
	if req.Details.Set {
		details := req.Details.Value
		if details.DiagnosticMarking.Set {
			marking := domain.DiagnosticMarking(details.DiagnosticMarking.Value)
			arg.DiagnosticMarking = &marking
		}
		if details.MaterialType.Set {
			materialType := domain.MaterialType(details.MaterialType.Value)
			arg.MaterialType = &materialType
		}
		if details.Calcitonin.Set {
			calcitonin := int(details.Calcitonin.Value)
			arg.Calcitonin = &calcitonin
		}
		if details.CalcitoninInFlush.Set {
			calcitoninInFlush := int(details.CalcitoninInFlush.Value)
			arg.CalcitoninInFlush = &calcitoninInFlush
		}
		if details.Thyroglobulin.Set {
			thyroglobulin := int(details.Thyroglobulin.Value)
			arg.Thyroglobulin = &thyroglobulin
		}
	}

	// Также проверяем верхний уровень
	if req.DiagnosticMarking.Set {
		marking := domain.DiagnosticMarking(req.DiagnosticMarking.Value)
		arg.DiagnosticMarking = &marking
	}
	if req.MaterialType.Set {
		materialType := domain.MaterialType(req.MaterialType.Value)
		arg.MaterialType = &materialType
	}
	if req.Calcitonin.Set {
		calcitonin := int(req.Calcitonin.Value)
		arg.Calcitonin = &calcitonin
	}
	if req.CalcitoninInFlush.Set {
		calcitoninInFlush := int(req.CalcitoninInFlush.Value)
		arg.CalcitoninInFlush = &calcitoninInFlush
	}
	if req.Thyroglobulin.Set {
		thyroglobulin := int(req.Thyroglobulin.Value)
		arg.Thyroglobulin = &thyroglobulin
	}
	if req.IsLast.Set {
		arg.IsLast = &req.IsLast.Value
	}

	// Обработка prev и parent_prev из req.Details
	if req.Details.Set {
		details := req.Details.Value
		if details.Prev.Set {
			arg.PrevID = &details.Prev.Value
		}
		if details.ParentPrev.Set {
			arg.ParentPrevID = &details.ParentPrev.Value
		}
	}

	// Также проверяем верхний уровень
	if req.Prev.Set {
		arg.PrevID = &req.Prev.Value
	}
	if req.ParentPrev.Set {
		arg.ParentPrevID = &req.ParentPrev.Value
	}

	return arg
}

func (CytologyImage) UpdateArgFromCytologyUpdatePartialUpdateReq(id uuid.UUID, req *api.CytologyUpdatePartialUpdateReq) cytologySrv.UpdateCytologyImageArg {
	arg := cytologySrv.UpdateCytologyImageArg{
		Id: id,
	}

	if req.DiagnosticMarking.Set {
		marking := domain.DiagnosticMarking(req.DiagnosticMarking.Value)
		arg.DiagnosticMarking = &marking
	}
	if req.MaterialType.Set {
		materialType := domain.MaterialType(req.MaterialType.Value)
		arg.MaterialType = &materialType
	}
	if req.Calcitonin.Set {
		calcitonin := int(req.Calcitonin.Value)
		arg.Calcitonin = &calcitonin
	}
	if req.CalcitoninInFlush.Set {
		calcitoninInFlush := int(req.CalcitoninInFlush.Value)
		arg.CalcitoninInFlush = &calcitoninInFlush
	}
	if req.Thyroglobulin.Set {
		thyroglobulin := int(req.Thyroglobulin.Value)
		arg.Thyroglobulin = &thyroglobulin
	}
	if req.IsLast.Set {
		arg.IsLast = &req.IsLast.Value
	}

	// Обработка prev и parent_prev
	if req.Prev.Set {
		arg.PrevID = &req.Prev.Value
	}
	if req.ParentPrev.Set {
		arg.ParentPrevID = &req.ParentPrev.Value
	}

	return arg
}

func (CytologyImage) ToCytologyUpdateUpdateOK(img domain.CytologyImage, req *api.CytologyUpdateUpdateReq) api.CytologyUpdateUpdateOK {
	result := api.CytologyUpdateUpdateOK{
		PatientCard: api.CytologyUpdateUpdateOKPatientCard{},
		IsLast: api.OptBool{
			Value: img.IsLast,
			Set:   true,
		},
		DiagnosDate: api.OptDateTime{
			Value: img.DiagnosDate,
			Set:   true,
		},
		DiagnosticNumber: int(img.DiagnosticNumber),
	}

	if img.DiagnosticMarking != nil {
		result.DiagnosticMarking = api.OptString{
			Value: string(*img.DiagnosticMarking),
			Set:   true,
		}
	}

	if img.MaterialType != nil {
		result.MaterialType = api.OptString{
			Value: string(*img.MaterialType),
			Set:   true,
		}
	}

	if img.Calcitonin != nil {
		result.Calcitonin = api.OptInt{
			Value: int(*img.Calcitonin),
			Set:   true,
		}
	}

	if img.CalcitoninInFlush != nil {
		result.CalcitoninInFlush = api.OptInt{
			Value: int(*img.CalcitoninInFlush),
			Set:   true,
		}
	}

	if img.Thyroglobulin != nil {
		result.Thyroglobulin = api.OptInt{
			Value: int(*img.Thyroglobulin),
			Set:   true,
		}
	}

	if img.Details != nil {
		result.Details = &api.CytologyUpdateUpdateOKDetails{}
	}

	// Маппинг prev и parent_prev
	if img.PrevID != nil {
		result.Prev = api.OptUUID{
			Value: *img.PrevID,
			Set:   true,
		}
	}
	if img.ParentPrevID != nil {
		result.ParentPrev = api.OptUUID{
			Value: *img.ParentPrevID,
			Set:   true,
		}
	}

	return result
}

func (CytologyImage) ToCytologyUpdatePartialUpdateOK(img domain.CytologyImage, req *api.CytologyUpdatePartialUpdateReq) api.CytologyUpdatePartialUpdateOK {
	result := api.CytologyUpdatePartialUpdateOK{
		PatientCard: api.CytologyUpdatePartialUpdateOKPatientCard{},
		IsLast: api.OptBool{
			Value: img.IsLast,
			Set:   true,
		},
		DiagnosDate: api.OptDateTime{
			Value: img.DiagnosDate,
			Set:   true,
		},
		DiagnosticNumber: int(img.DiagnosticNumber),
	}

	if img.DiagnosticMarking != nil {
		result.DiagnosticMarking = api.OptString{
			Value: string(*img.DiagnosticMarking),
			Set:   true,
		}
	}

	if img.MaterialType != nil {
		result.MaterialType = api.OptString{
			Value: string(*img.MaterialType),
			Set:   true,
		}
	}

	if img.Calcitonin != nil {
		result.Calcitonin = api.OptInt{
			Value: int(*img.Calcitonin),
			Set:   true,
		}
	}

	if img.CalcitoninInFlush != nil {
		result.CalcitoninInFlush = api.OptInt{
			Value: int(*img.CalcitoninInFlush),
			Set:   true,
		}
	}

	if img.Thyroglobulin != nil {
		result.Thyroglobulin = api.OptInt{
			Value: int(*img.Thyroglobulin),
			Set:   true,
		}
	}

	if img.Details != nil {
		result.Details = &api.CytologyUpdatePartialUpdateOKDetails{}
	}

	// Маппинг prev и parent_prev
	if img.PrevID != nil {
		result.Prev = api.OptUUID{
			Value: *img.PrevID,
			Set:   true,
		}
	}
	if img.ParentPrevID != nil {
		result.ParentPrev = api.OptUUID{
			Value: *img.ParentPrevID,
			Set:   true,
		}
	}

	return result
}
