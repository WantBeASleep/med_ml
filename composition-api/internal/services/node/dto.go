package node

import (
	domain "composition-api/internal/domain/uzi"

	"github.com/google/uuid"
)

type UpdateNodeArg struct {
	Id         uuid.UUID
	Validation *domain.NodeValidation
	Tirads_23  *float64
	Tirads_4   *float64
	Tirads_5   *float64
}
