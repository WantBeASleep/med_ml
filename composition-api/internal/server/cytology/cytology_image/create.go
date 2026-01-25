package cytology_image

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"composition-api/internal/domain"

	"github.com/AlekSi/pointer"

	api "composition-api/internal/generated/http/api"
	mappers "composition-api/internal/server/cytology/mappers"
)

func (h *handler) CytologyCreateCreate(ctx context.Context, req *api.CytologyCreateCreateReq) (api.CytologyCreateCreateRes, error) {
	slog.Info("CytologyCreateCreate: received request",
		"diagnostic_number", req.DiagnosticNumber,
		"has_image", req.Image.Set,
		"image_size", func() int64 {
			if req.Image.Set {
				return req.Image.Value.Size
			}
			return 0
		}(),
	)

	arg := mappers.CytologyImage{}.CreateArgFromCytologyCreateCreateReq(req)

	startTime := time.Now()
	id, err := h.services.CytologyService.CreateCytologyImage(ctx, arg)
	duration := time.Since(startTime)

	if err != nil {
		slog.Error("CytologyCreateCreate: failed to create",
			"err", err,
			"duration_ms", duration.Milliseconds(),
		)
		switch {
		case errors.Is(err, domain.ErrBadRequest):
			return &api.CytologyCreateCreateBadRequest{
				StatusCode: http.StatusBadRequest,
				Response: api.Error{
					Message: "Неверный формат запроса",
				},
			}, nil
		case errors.Is(err, domain.ErrUnprocessableEntity):
			return &api.CytologyCreateCreateUnprocessableEntity{
				StatusCode: http.StatusUnprocessableEntity,
				Response: api.Error{
					Message: "Ошибка валидации данных",
				},
			}, nil
		default:
			return nil, err
		}
	}

	// Получаем созданное исследование для заполнения всех полей ответа
	img, origImg, err := h.services.CytologyService.GetCytologyImageById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Возвращаем созданный объект согласно swagger.json
	result := api.CytologyCreateCreateCreated{
		DiagnosticNumber: req.DiagnosticNumber,
		ID: api.OptUUID{
			Value: id,
			Set:   true,
		},
	}

	// Заполняем поля из созданного объекта
	if origImg != nil && origImg.ImagePath != "" {
		imageURLStr := id.String() + "/" + origImg.Id.String()
		imageURL, err := url.Parse(imageURLStr)
		if err == nil {
			result.Image = api.OptURI{
				Value: *imageURL,
				Set:   true,
			}
		}
	}

	result.IsLast = api.OptBool{
		Value: img.IsLast,
		Set:   true,
	}

	result.DiagnosDate = api.OptDateTime{
		Value: img.DiagnosDate,
		Set:   true,
	}

	if req.Details.Set {
		result.Details = &api.CytologyCreateCreateCreatedDetails{}
	}
	if req.DiagnosticMarking.Set {
		result.DiagnosticMarking = api.OptCytologyCreateCreateCreatedDiagnosticMarking{
			Value: api.CytologyCreateCreateCreatedDiagnosticMarking(req.DiagnosticMarking.Value),
			Set:   true,
		}
	}
	if req.MaterialType.Set {
		result.MaterialType = api.OptCytologyCreateCreateCreatedMaterialType{
			Value: api.CytologyCreateCreateCreatedMaterialType(req.MaterialType.Value),
			Set:   true,
		}
	}
	if req.Calcitonin.Set {
		result.Calcitonin = req.Calcitonin
	}
	if req.CalcitoninInFlush.Set {
		result.CalcitoninInFlush = req.CalcitoninInFlush
	}
	if req.Thyroglobulin.Set {
		result.Thyroglobulin = req.Thyroglobulin
	}
	if req.Prev.Set {
		result.Prev = req.Prev
	}
	if req.ParentPrev.Set {
		result.ParentPrev = req.ParentPrev
	}
	if req.PatientCard.Set {
		result.PatientCard = req.PatientCard
	}

	slog.Info("CytologyCreateCreate: successfully created",
		"id", id,
		"duration_ms", duration.Milliseconds(),
	)

	return pointer.To(result), nil
}
