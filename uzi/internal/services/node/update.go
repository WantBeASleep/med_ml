package node

import (
	"context"
	"fmt"

	"uzi/internal/domain"
	nodeEntity "uzi/internal/repository/node/entity"
)

func (s *service) UpdateNode(ctx context.Context, arg UpdateNodeArg) (domain.Node, error) {
	nodeQuery := s.dao.NewNodeQuery(ctx)

	nodeDB, err := nodeQuery.GetNodeByID(arg.Id)
	if err != nil {
		return domain.Node{}, fmt.Errorf("get node: %w", err)
	}
	node := nodeDB.ToDomain()
	arg.UpdateDomain(&node)

	if err := nodeQuery.UpdateNode(nodeEntity.Node{}.FromDomain(node)); err != nil {
		return domain.Node{}, fmt.Errorf("update node: %w", err)
	}

	return node, nil
}
