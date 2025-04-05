package uzi

import (
	"context"

	"github.com/AlekSi/pointer"

	api "composition-api/internal/generated/http/api"
	mappers "composition-api/internal/server/uzi/mappers"
)

func (h *handler) UziIDGet(ctx context.Context, params api.UziIDGetParams) (api.UziIDGetRes, error) {
	uzi, err := h.services.UziService.GetByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	return pointer.To(mappers.Uzi{}.Domain(uzi)), nil
}

func (h *handler) UzisExternalIDGet(ctx context.Context, params api.UzisExternalIDGetParams) (api.UzisExternalIDGetRes, error) {
	uzis, err := h.services.UziService.GetByExternalID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	return pointer.To(api.UzisExternalIDGetOKApplicationJSON(mappers.Uzi{}.SliceDomain(uzis))), nil
}

func (h *handler) UzisAuthorIDGet(ctx context.Context, params api.UzisAuthorIDGetParams) (api.UzisAuthorIDGetRes, error) {
	uzis, err := h.services.UziService.GetByAuthor(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	return pointer.To(api.UzisAuthorIDGetOKApplicationJSON(mappers.Uzi{}.SliceDomain(uzis))), nil
}

func (h *handler) UziIDEchographicsGet(ctx context.Context, params api.UziIDEchographicsGetParams) (api.UziIDEchographicsGetRes, error) {
	echographics, err := h.services.UziService.GetEchographicsByID(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	return pointer.To(mappers.Echographics(echographics)), nil
}
