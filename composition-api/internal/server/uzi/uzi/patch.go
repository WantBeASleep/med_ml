package uzi

import (
	"context"

	"github.com/AlekSi/pointer"

	domain "composition-api/internal/domain/uzi"
	api "composition-api/internal/generated/http/api"
	apimappers "composition-api/internal/server/mappers"
	mappers "composition-api/internal/server/uzi/mappers"
	uziSrv "composition-api/internal/services/uzi"
)

func (h *handler) UziIDPatch(ctx context.Context, req *api.UziIDPatchReq, params api.UziIDPatchParams) (api.UziIDPatchRes, error) {
	var projection *domain.UziProjection
	if req.Projection.IsSet() {
		projection = (*domain.UziProjection)(&req.Projection.Value)
	}

	uzi, err := h.services.UziService.Update(ctx, uziSrv.UpdateUziArg{
		Id:         params.ID,
		Projection: projection,
		Checked:    apimappers.FromOptBool(req.Checked),
	})
	if err != nil {
		return nil, err
	}
	return pointer.To(mappers.Uzi{}.Domain(uzi)), nil
}

func (h *handler) UziIDEchographicsPatch(ctx context.Context, req *api.Echographics, params api.UziIDEchographicsPatchParams) (api.UziIDEchographicsPatchRes, error) {
	echographics, err := h.services.UziService.UpdateEchographics(ctx, domain.Echographic{
		Id:              params.ID,
		Contors:         apimappers.FromOptString(req.Contors),
		LeftLobeLength:  apimappers.FromOptFloat64(req.LeftLobeLength),
		LeftLobeWidth:   apimappers.FromOptFloat64(req.LeftLobeWidth),
		LeftLobeThick:   apimappers.FromOptFloat64(req.LeftLobeThick),
		LeftLobeVolum:   apimappers.FromOptFloat64(req.LeftLobeVolum),
		RightLobeLength: apimappers.FromOptFloat64(req.RightLobeLength),
		RightLobeWidth:  apimappers.FromOptFloat64(req.RightLobeWidth),
		RightLobeThick:  apimappers.FromOptFloat64(req.RightLobeThick),
		RightLobeVolum:  apimappers.FromOptFloat64(req.RightLobeVolum),
		GlandVolum:      apimappers.FromOptFloat64(req.GlandVolum),
		Isthmus:         apimappers.FromOptFloat64(req.Isthmus),
		Struct:          apimappers.FromOptString(req.Struct),
		Echogenicity:    apimappers.FromOptString(req.Echogenicity),
		RegionalLymph:   apimappers.FromOptString(req.RegionalLymph),
		Vascularization: apimappers.FromOptString(req.Vascularization),
		Location:        apimappers.FromOptString(req.Location),
		Additional:      apimappers.FromOptString(req.Additional),
		Conclusion:      apimappers.FromOptString(req.Conclusion),
	})
	if err != nil {
		return nil, err
	}
	return pointer.To(mappers.Echographics(echographics)), nil
}
