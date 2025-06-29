package refresh_token

import (
	daolib "github.com/WantBeASleep/med_ml_lib/dao"
	"github.com/google/uuid"

	rtentity "auth/internal/repository/refresh_token/entity"
)

const (
	table              = "refresh_token"
	columnID           = "id"
	columnRefreshToken = "refresh_token"
)

type Repository interface {
	InsertRefreshToken(refreshToken rtentity.RefreshToken) error
	ExistsRefreshToken(id uuid.UUID, refreshToken string) (bool, error)
	DeleteRefreshTokens([]rtentity.RefreshToken) error
}

type repo struct {
	*daolib.BaseQuery
}

func NewR() *repo {
	return &repo{}
}

func (r *repo) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	r.BaseQuery = baseQuery
}
