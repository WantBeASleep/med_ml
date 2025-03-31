package repository

import (
	"context"

	"auth/internal/repository/user"

	daolib "github.com/WantBeASleep/med_ml_lib/dao"

	"github.com/jmoiron/sqlx"
)

type DAO interface {
	daolib.DAO
	NewUserRepo(ctx context.Context) user.Repository
}

type dao struct {
	daolib.DAO
}

func NewRepository(psql *sqlx.DB) DAO {
	return &dao{DAO: daolib.NewDao(psql)}
}

func (d *dao) NewUserRepo(ctx context.Context) user.Repository {
	userRepo := user.NewR()
	d.NewRepo(ctx, userRepo)

	return userRepo
}
