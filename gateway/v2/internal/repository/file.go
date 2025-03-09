package repository

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime"
	"path/filepath"

	ht "github.com/ogen-go/ogen/http"

	minio "github.com/minio/minio-go/v7"
)

var ErrFileNotFound = errors.New("file not found")

type FileRepo interface {
	GetFile(ctx context.Context, path string) (io.ReadCloser, error)
	LoadFile(ctx context.Context, path string, file ht.MultipartFile) error
}

type fileRepo struct {
	s3     *minio.Client
	bucket string
}

func (r *fileRepo) GetFile(ctx context.Context, path string) (io.ReadCloser, error) {
	_, err := r.s3.StatObject(ctx, r.bucket, path, minio.GetObjectOptions{})
	if err != nil {
		errResp := minio.ToErrorResponse(err)
		if errResp.Code == "NoSuchKey" {
			return nil, ErrFileNotFound
		}
		return nil, fmt.Errorf("get stat of object: %w", err)
	}

	obj, err := r.s3.GetObject(ctx, r.bucket, path, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (r *fileRepo) LoadFile(ctx context.Context, path string, file ht.MultipartFile) error {
	_, err := r.s3.PutObject(ctx, r.bucket, path, file.File, -1, minio.PutObjectOptions{
		ContentType: mime.TypeByExtension(filepath.Ext(file.Name)),
	})
	if err != nil {
		return err
	}

	return nil
}
