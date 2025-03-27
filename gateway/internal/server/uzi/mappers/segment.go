package mappers

import (
	domain "composition-api/internal/domain/uzi"
	api "composition-api/internal/generated/http/api"
)

func Segment(segment domain.Segment) (api.Segment, error) {
	contor, err := Contor(segment.Contor)
	if err != nil {
		return api.Segment{}, err
	}

	return api.Segment{
		ID:       segment.Id,
		ImageID:  segment.ImageID,
		NodeID:   segment.NodeID,
		Contor:   contor,
		Tirads23: segment.Tirads23,
		Tirads4:  segment.Tirads4,
		Tirads5:  segment.Tirads5,
	}, nil
}

func SliceSegment(segments []domain.Segment) ([]api.Segment, error) {
	result := make([]api.Segment, 0, len(segments))
	for _, segment := range segments {
		segment, err := Segment(segment)
		if err != nil {
			return nil, err
		}
		result = append(result, segment)
	}
	return result, nil
}
