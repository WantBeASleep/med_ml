package repository

import (
	"context"

	daolib "github.com/WantBeASleep/med_ml_lib/dao"
	"github.com/jmoiron/sqlx"
	minio "github.com/minio/minio-go/v7"

	"cytology/internal/repository/cytology_image"
	"cytology/internal/repository/original_image"
	"cytology/internal/repository/segmentation"
	"cytology/internal/repository/segmentation_group"
)

type DAO interface {
	daolib.DAO
	NewFileRepo() FileRepo
	NewCytologyImageQuery(ctx context.Context) cytology_image.Repository
	NewOriginalImageQuery(ctx context.Context) original_image.Repository
	NewSegmentationGroupQuery(ctx context.Context) segmentation_group.Repository
	NewSegmentationQuery(ctx context.Context) segmentation.Repository
}

type dao struct {
	daolib.DAO

	s3       *minio.Client
	s3bucket string
}

func NewRepository(psql *sqlx.DB, s3 *minio.Client, s3bucket string) DAO {
	return &dao{
		DAO:      daolib.NewDao(psql),
		s3:       s3,
		s3bucket: s3bucket,
	}
}

// S3
func (d *dao) NewFileRepo() FileRepo {
	return &fileRepo{
		s3:     d.s3,
		bucket: d.s3bucket,
	}
}

// POSTGRES
func (d *dao) NewCytologyImageQuery(ctx context.Context) cytology_image.Repository {
	query := cytology_image.NewR()
	d.NewRepo(ctx, query)
	return query
}

func (d *dao) NewOriginalImageQuery(ctx context.Context) original_image.Repository {
	query := original_image.NewR()
	d.NewRepo(ctx, query)
	return query
}

func (d *dao) NewSegmentationGroupQuery(ctx context.Context) segmentation_group.Repository {
	query := segmentation_group.NewR()
	d.NewRepo(ctx, query)
	return query
}

func (d *dao) NewSegmentationQuery(ctx context.Context) segmentation.Repository {
	query := segmentation.NewR()
	d.NewRepo(ctx, query)
	return query
}
