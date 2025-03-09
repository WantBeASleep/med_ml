package echographic

import (
	"uzi/internal/repository/echographic/entity"

	daolib "github.com/WantBeASleep/med_ml_lib/dao"
	"github.com/google/uuid"
)

const (
	table = "echographic"

	columnID              = "id"
	columnContors         = "contors"
	columnLeftLobeLength  = "left_lobe_length"
	columnLeftLobeWidth   = "left_lobe_width"
	columnLeftLobeThick   = "left_lobe_thick"
	columnLeftLobeVolum   = "left_lobe_volum"
	columnRightLobeLength = "right_lobe_length"
	columnRightLobeWidth  = "right_lobe_width"
	columnRightLobeThick  = "right_lobe_thick"
	columnRightLobeVolum  = "right_lobe_volum"
	columnGlandVolum      = "gland_volum"
	columnIsthmus         = "isthmus"
	columnStruct          = "struct"
	columnEchogenicity    = "echogenicity"
	columnRegionalLymph   = "regional_lymph"
	columnVascularization = "vascularization"
	columnLocation        = "location"
	columnAdditional      = "additional"
	columnConclusion      = "conclusion"
)

type Repository interface {
	InsertEchographic(echographic entity.Echographic) error

	GetEchographicByID(id uuid.UUID) (entity.Echographic, error)

	UpdateEchographic(echographic entity.Echographic) error
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
