package mappers

import (
	domain "gateway/internal/domain/uzi"
	api "gateway/internal/generated/http/api"
	mappers "gateway/internal/server/mappers"
)

func Echographics(echographics domain.Echographic) api.Echographics {
	return api.Echographics{
		ID:              echographics.Id,
		Contors:         mappers.ToOptString(echographics.Contors),
		LeftLobeLength:  mappers.ToOptFloat64(echographics.LeftLobeLength),
		LeftLobeWidth:   mappers.ToOptFloat64(echographics.LeftLobeWidth),
		LeftLobeThick:   mappers.ToOptFloat64(echographics.LeftLobeThick),
		LeftLobeVolum:   mappers.ToOptFloat64(echographics.LeftLobeVolum),
		RightLobeLength: mappers.ToOptFloat64(echographics.RightLobeLength),
		RightLobeWidth:  mappers.ToOptFloat64(echographics.RightLobeWidth),
		RightLobeThick:  mappers.ToOptFloat64(echographics.RightLobeThick),
		RightLobeVolum:  mappers.ToOptFloat64(echographics.RightLobeVolum),
		GlandVolum:      mappers.ToOptFloat64(echographics.GlandVolum),
		Isthmus:         mappers.ToOptFloat64(echographics.Isthmus),
		Struct:          mappers.ToOptString(echographics.Struct),
		Echogenicity:    mappers.ToOptString(echographics.Echogenicity),
		RegionalLymph:   mappers.ToOptString(echographics.RegionalLymph),
		Vascularization: mappers.ToOptString(echographics.Vascularization),
		Location:        mappers.ToOptString(echographics.Location),
		Additional:      mappers.ToOptString(echographics.Additional),
		Conclusion:      mappers.ToOptString(echographics.Conclusion),
	}
}

func SliceEchographics(echographics []domain.Echographic) []api.Echographics {
	result := make([]api.Echographics, 0, len(echographics))
	for _, echographic := range echographics {
		result = append(result, Echographics(echographic))
	}
	return result
}
