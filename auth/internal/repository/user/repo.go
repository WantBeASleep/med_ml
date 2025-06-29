package user

import (
	daolib "github.com/WantBeASleep/med_ml_lib/dao"

	uentity "auth/internal/repository/user/entity"

	"github.com/google/uuid"
)

const (
	table          = "\"user\""
	columnID       = "id"
	columnEmail    = "email"
	columnPassword = "password"
	columnRole     = "role"
)

type Repository interface {
	InsertUser(user uentity.User) error

	GetUserByID(id uuid.UUID) (uentity.User, error)
	GetUserByEmail(email string) (uentity.User, error)

	UpdateUserPassword(id uuid.UUID, password string) error
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
