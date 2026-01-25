package cytology_image

import (
	"context"
	"errors"
	"net/http"

	"composition-api/internal/domain"
	med_domain "composition-api/internal/domain/med"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"

	api "composition-api/internal/generated/http/api"
	mappers "composition-api/internal/server/cytology/mappers"
)

func (h *handler) CytologyRead(ctx context.Context, params api.CytologyReadParams) (api.CytologyReadRes, error) {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return &api.CytologyReadInternalServerError{
			StatusCode: http.StatusBadRequest,
			Response: api.Error{
				Message: "Неверный формат ID",
			},
		}, nil
	}

	img, origImg, err := h.services.CytologyService.GetCytologyImageById(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return &api.CytologyReadNotFound{
				StatusCode: http.StatusNotFound,
				Response: api.Error{
					Message: "Цитологическое исследование не найдено",
				},
			}, nil
		default:
			return nil, err
		}
	}

	// Получаем данные о пациенте
	var patient med_domain.Patient
	if img.PatientID != uuid.Nil {
		patient, err = h.services.PatientService.GetPatient(ctx, img.PatientID)
		if err != nil {
			// Если не удалось получить пациента, продолжаем с пустым объектом
			patient = med_domain.Patient{}
		}
	}

	// Получаем данные о карточке пациента
	var patientCard med_domain.Card
	if img.DoctorID != uuid.Nil && img.PatientID != uuid.Nil {
		patientCard, err = h.services.CardService.GetCard(ctx, img.DoctorID, img.PatientID)
		if err != nil {
			// Если не удалось получить карточку, продолжаем с пустым объектом
			patientCard = med_domain.Card{}
		}
	}

	// Маппим в структуру согласно swagger.json
	result := api.CytologyReadOK{
		OriginalImage: mappers.OriginalImage{}.ToCytologyReadOKOriginalImage(origImg),
		Info:          mappers.CytologyImage{}.ToCytologyReadOKInfo(img, patient, patientCard),
	}

	return pointer.To(result), nil
}
