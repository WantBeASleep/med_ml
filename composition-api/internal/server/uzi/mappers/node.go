package mappers

import (
	domain "composition-api/internal/domain/uzi"
	api "composition-api/internal/generated/http/api"
	apimappers "composition-api/internal/server/mappers"
)

type Node struct{}

func (Node) Domain(node domain.Node) api.Node {
	var validation api.OptNilNodeValidation
	if node.Validation != nil {
		validation.Set = true

		switch *node.Validation {
		case domain.NodeValidationNull:
			validation.Null = true
		case domain.NodeValidationInvalid:
			validation.Value = api.NodeValidationInvalid
		case domain.NodeValidationValid:
			validation.Value = api.NodeValidationValid
		}
	}

	return api.Node{
		ID:          node.Id,
		Ai:          node.Ai,
		UziID:       node.UziID,
		Validation:  validation,
		Tirads23:    node.Tirads23,
		Tirads4:     node.Tirads4,
		Tirads5:     node.Tirads5,
		Description: apimappers.ToOptString(node.Description),
	}
}

func (Node) SliceDomain(nodes []domain.Node) []api.Node {
	return slice(nodes, Node{})
}
