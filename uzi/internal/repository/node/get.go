package node

import (
	"database/sql"
	"errors"
	"fmt"

	daoEntity "uzi/internal/repository/entity"
	"uzi/internal/repository/node/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (q *repo) GetNodeByID(id uuid.UUID) (entity.Node, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnAI,
			columnUziID,
			columnTirads23,
			columnTirads4,
			columnTirads5,
		).
		From(table).
		Where(sq.Eq{
			columnID: id,
		})

	var node entity.Node
	if err := q.Runner().Getx(q.Context(), &node, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Node{}, daoEntity.ErrNotFound
		}
		return entity.Node{}, err
	}

	return node, nil
}

func (q *repo) GetNodesByImageID(id uuid.UUID) ([]entity.Node, error) {
	query := q.QueryBuilder().
		Select(
			fmt.Sprintf("%s.%s", table, columnID),
			fmt.Sprintf("%s.%s", table, columnAI),
			fmt.Sprintf("%s.%s", table, columnUziID),
			fmt.Sprintf("%s.%s", table, columnTirads23),
			fmt.Sprintf("%s.%s", table, columnTirads4),
			fmt.Sprintf("%s.%s", table, columnTirads5),
		).
		From(table). // TODO: вынести константы таблиц в отдельный пакет, тут пересечение с segment
		InnerJoin("segment ON segment.node_id = node.id").
		InnerJoin("image ON image.id = segment.image_id").
		Where(sq.Eq{
			"image.id": id,
		})

	var uzi []entity.Node
	if err := q.Runner().Selectx(q.Context(), &uzi, query); err != nil {
		return nil, err
	}

	if len(uzi) == 0 {
		return nil, daoEntity.ErrNotFound
	}

	return uzi, nil
}

func (q *repo) GetNodesByUziID(id uuid.UUID) ([]entity.Node, error) {
	query := q.QueryBuilder().
		Select(
			columnID,
			columnAI,
			columnUziID,
			columnTirads23,
			columnTirads4,
			columnTirads5,
		).
		From(table).
		Where(sq.Eq{
			columnUziID: id,
		})

	var uzi []entity.Node
	if err := q.Runner().Selectx(q.Context(), &uzi, query); err != nil {
		return nil, err
	}

	if len(uzi) == 0 {
		return nil, daoEntity.ErrNotFound
	}

	return uzi, nil
}
