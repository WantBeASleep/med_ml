package mappers

import (
	domain "gateway/internal/domain/uzi"
	api "gateway/internal/generated/http/api"
)

func Node(node domain.Node) api.Node {
	return api.Node{
		ID:       node.Id,
		Ai:       node.Ai,
		UziID:    node.UziID,
		Tirads23: node.Tirads23,
		Tirads4:  node.Tirads4,
		Tirads5:  node.Tirads5,
	}
}

func SliceNode(nodes []domain.Node) []api.Node {
	result := make([]api.Node, 0, len(nodes))
	for _, node := range nodes {
		result = append(result, Node(node))
	}
	return result
}
