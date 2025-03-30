package node

import (
	"context"
	"errors"
	"fmt"

	"uzi/internal/domain"
	nodeEntity "uzi/internal/repository/node/entity"
)

var (
	ErrAiNodeEdit = errors.New("unable to edit ai node")
)

func (s *service) UpdateNode(ctx context.Context, arg UpdateNodeArg) (domain.Node, error) {
	nodeQuery := s.dao.NewNodeQuery(ctx)

	nodeDB, err := nodeQuery.GetNodeByID(arg.Id)
	if err != nil {
		return domain.Node{}, fmt.Errorf("get node: %w", err)
	}
	node := nodeDB.ToDomain()

	// валидация по AI
	switch node.Ai {
	case true:
		if arg.Tirads23 != nil ||
			arg.Tirads4 != nil ||
			arg.Tirads5 != nil {
			return domain.Node{}, ErrAiNodeEdit
		}

		if arg.Validation == nil {
			return domain.Node{}, fmt.Errorf("can't set nil validation for ai node")
		}

	case false:
		if arg.Validation != nil {
			return domain.Node{}, fmt.Errorf("can't set validation for manual node")
		}
	}

	arg.UpdateDomain(&node)
	if err := nodeQuery.UpdateNode(nodeEntity.Node{}.FromDomain(node)); err != nil {
		return domain.Node{}, fmt.Errorf("update node: %w", err)
	}

	return node, nil
}
