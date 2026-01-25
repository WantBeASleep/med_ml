package mappers

import (
	"encoding/json"

	"github.com/google/uuid"

	domain "composition-api/internal/domain/cytology"
	api "composition-api/internal/generated/http/api"
	cytologySrv "composition-api/internal/services/cytology"
)

type SegmentationGroup struct{}

// Удалены неиспользуемые методы Domain, SliceDomain, CreateArg, UpdateArg - заменены на новые методы для работы с обновленными типами API

type Segmentation struct{}

// Удалены неиспользуемые методы Domain, SliceDomain, CreateArg, UpdateArg - заменены на новые методы для работы с обновленными типами API

func (SegmentationGroup) ToSegmentationDataList(groups []domain.SegmentationGroup, getSegments func(int) ([]domain.Segmentation, error)) []api.CytologySegmentsListOKResultsItem {
	result := make([]api.CytologySegmentsListOKResultsItem, 0, len(groups))
	for _, group := range groups {
		// Получаем сегменты для группы
		segments, err := getSegments(group.Id)
		if err != nil {
			// Если ошибка, продолжаем с пустым списком сегментов
			segments = []domain.Segmentation{}
		}

		// Маппим сегменты в формат данных
		dataItems := make([]api.CytologySegmentsListOKResultsItemDataItem, 0, len(segments))
		for _, seg := range segments {
			points := make([]api.CytologySegmentsListOKResultsItemDataItemPointsItem, 0, len(seg.Points))
			for _, point := range seg.Points {
				points = append(points, api.CytologySegmentsListOKResultsItemDataItemPointsItem{
					ID: api.OptInt{
						Value: point.Id,
						Set:   true,
					},
					UID: int(point.UID), // UID имеет тип int в API
					X:   point.X,
					Y:   point.Y,
				})
			}
			dataItems = append(dataItems, api.CytologySegmentsListOKResultsItemDataItem{
				ID: api.OptInt{
					Value: seg.Id,
					Set:   true,
				},
				Points: points,
				Details: api.OptString{
					// Details может быть в группе, но не в сегменте
					Set: false,
				},
			})
		}

		item := api.CytologySegmentsListOKResultsItem{
			ID: api.OptInt{
				Value: group.Id,
				Set:   true,
			},
			Data: dataItems,
			GroupType: api.OptCytologySegmentsListOKResultsItemGroupType{
				Value: api.CytologySegmentsListOKResultsItemGroupType(group.GroupType),
				Set:   true,
			},
			SegType: api.OptCytologySegmentsListOKResultsItemSegType{
				Value: api.CytologySegmentsListOKResultsItemSegType(group.SegType),
				Set:   true,
			},
			IsAi: api.OptBool{
				Value: group.IsAI,
				Set:   true,
			},
		}

		// Маппим details из группы, если они есть
		if group.Details != nil {
			// Пытаемся распарсить JSON из строки
			var detailsObj map[string]interface{}
			if err := json.Unmarshal([]byte(*group.Details), &detailsObj); err == nil {
				item.Details = &api.CytologySegmentsListOKResultsItemDetails{}
			} else {
				item.Details = &api.CytologySegmentsListOKResultsItemDetails{}
			}
		}

		result = append(result, item)
	}
	return result
}

func (SegmentationGroup) CreateArgFromCytologySegmentGroupCreateCreateReq(cytologyID uuid.UUID, req *api.CytologySegmentGroupCreateCreateReq) (cytologySrv.CreateSegmentationGroupArg, cytologySrv.CreateSegmentationArg) {
	// Определяем group_type из seg_type
	segType := domain.SegType(req.SegType)
	groupType := domain.GroupTypeCE // По умолчанию
	switch segType {
	case domain.SegTypeNIL, domain.SegTypeCNO, domain.SegTypeCGE, domain.SegTypeC2N, domain.SegTypeCPS, domain.SegTypeCFC, domain.SegTypeCLY:
		groupType = domain.GroupTypeCE
	case domain.SegTypeNIR, domain.SegTypeSOS, domain.SegTypeSDS, domain.SegTypeSMS, domain.SegTypeSTS, domain.SegTypeSPS:
		groupType = domain.GroupTypeCL
	case domain.SegTypeNIM, domain.SegTypeSNM, domain.SegTypeSTM:
		groupType = domain.GroupTypeME
	}

	// Details не передаются в запросе, оставляем nil
	// Они могут быть установлены позже или через другую ручку
	groupArg := cytologySrv.CreateSegmentationGroupArg{
		CytologyID: cytologyID,
		SegType:    segType,
		GroupType:  groupType,
		IsAI:       false, // По умолчанию false для ручного создания
		Details:    nil,
	}

	// Создаем аргументы для создания сегмента с точками
	points := make([]domain.SegmentationPoint, 0, len(req.Data.Points))
	for _, point := range req.Data.Points {
		points = append(points, domain.SegmentationPoint{
			X: point.X,
			Y: point.Y,
		})
	}

	// Возвращаем оба аргумента - сначала нужно создать группу, потом сегмент
	// Но group ID будет известен только после создания, поэтому возвращаем только points
	segArg := cytologySrv.CreateSegmentationArg{
		SegmentationGroupID: 0, // Будет установлен после создания группы
		Points:              points,
	}

	return groupArg, segArg
}

func (SegmentationGroup) ToCytologySegmentGroupCreateCreateCreatedDataPoints(points []api.CytologySegmentGroupCreateCreateReqDataPointsItem) []api.CytologySegmentGroupCreateCreateCreatedDataPointsItem {
	result := make([]api.CytologySegmentGroupCreateCreateCreatedDataPointsItem, 0, len(points))
	for _, point := range points {
		result = append(result, api.CytologySegmentGroupCreateCreateCreatedDataPointsItem{
			X: point.X,
			Y: point.Y,
		})
	}
	return result
}

func (Segmentation) ToCytologySegmentUpdateReadOK(seg domain.Segmentation) api.CytologySegmentUpdateReadOK {
	result := api.CytologySegmentUpdateReadOK{
		Points: make([]api.CytologySegmentUpdateReadOKPointsItem, 0, len(seg.Points)),
		SegmentGroup: api.OptInt{
			Value: seg.SegmentationGroupID,
			Set:   true,
		},
	}

	for _, point := range seg.Points {
		result.Points = append(result.Points, api.CytologySegmentUpdateReadOKPointsItem{
			X: point.X,
			Y: point.Y,
		})
	}

	return result
}

func (Segmentation) UpdateArgFromCytologySegmentUpdateUpdateReq(id int, req *api.CytologySegmentUpdateUpdateReq) cytologySrv.UpdateSegmentationArg {
	points := make([]domain.SegmentationPoint, 0, len(req.Points))
	for _, point := range req.Points {
		points = append(points, domain.SegmentationPoint{
			X: point.X,
			Y: point.Y,
		})
	}

	return cytologySrv.UpdateSegmentationArg{
		Id:     id,
		Points: points,
	}
}

func (Segmentation) UpdateArgFromCytologySegmentUpdatePartialUpdateReq(id int, req *api.CytologySegmentUpdatePartialUpdateReq) cytologySrv.UpdateSegmentationArg {
	points := make([]domain.SegmentationPoint, 0, len(req.Points))
	for _, point := range req.Points {
		points = append(points, domain.SegmentationPoint{
			X: point.X,
			Y: point.Y,
		})
	}

	return cytologySrv.UpdateSegmentationArg{
		Id:     id,
		Points: points,
	}
}

func (Segmentation) ToCytologySegmentUpdateUpdateOK(seg domain.Segmentation, req *api.CytologySegmentUpdateUpdateReq) api.CytologySegmentUpdateUpdateOK {
	result := api.CytologySegmentUpdateUpdateOK{
		Points: make([]api.CytologySegmentUpdateUpdateOKPointsItem, 0, len(seg.Points)),
		SegmentGroup: api.OptInt{
			Value: seg.SegmentationGroupID,
			Set:   true,
		},
	}

	for _, point := range seg.Points {
		result.Points = append(result.Points, api.CytologySegmentUpdateUpdateOKPointsItem{
			X: point.X,
			Y: point.Y,
		})
	}

	return result
}

func (Segmentation) ToCytologySegmentUpdatePartialUpdateOK(seg domain.Segmentation, req *api.CytologySegmentUpdatePartialUpdateReq) api.CytologySegmentUpdatePartialUpdateOK {
	result := api.CytologySegmentUpdatePartialUpdateOK{
		Points: make([]api.CytologySegmentUpdatePartialUpdateOKPointsItem, 0, len(seg.Points)),
		SegmentGroup: api.OptInt{
			Value: seg.SegmentationGroupID,
			Set:   true,
		},
	}

	for _, point := range seg.Points {
		result.Points = append(result.Points, api.CytologySegmentUpdatePartialUpdateOKPointsItem{
			X: point.X,
			Y: point.Y,
		})
	}

	return result
}
