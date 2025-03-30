package node

import (
	"uzi/internal/repository/node/entity"

	daolib "github.com/WantBeASleep/med_ml_lib/dao"
	"github.com/google/uuid"
)

const (
	table = "node"

	columnID         = "id"
	columnAI         = "ai"
	columnUziID      = "uzi_id"
	columnValidation = "validation"
	columnTirads23   = "tirads_23"
	columnTirads4    = "tirads_4"
	columnTirads5    = "tirads_5"
)

type Repository interface {
	InsertNodes(nodes ...entity.Node) error

	GetNodeByID(id uuid.UUID) (entity.Node, error)
	GetNodesByImageID(id uuid.UUID) ([]entity.Node, error)
	GetNodesByUziID(id uuid.UUID) ([]entity.Node, error)

	UpdateNode(node entity.Node) error

	DeleteNodeByID(id uuid.UUID) error
}

type repo struct {
	*daolib.BaseQuery
}

func NewRepo() *repo {
	return &repo{}
}

func (q *repo) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}
