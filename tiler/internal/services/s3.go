package services

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/minio/minio-go/v7"
)

type s3Client struct {
	client     *minio.Client
	bucketName string
}

func NewS3Client(client *minio.Client, bucketName string) S3Client {
	return &s3Client{
		client:     client,
		bucketName: bucketName,
	}
}

func (s *s3Client) GetObject(ctx context.Context, bucketName, objectName string) ([]byte, error) {
	if bucketName == "" {
		bucketName = s.bucketName
	}

	// Убираем ведущий слэш из objectName, если он есть
	objectName = strings.TrimPrefix(objectName, "/")

	obj, err := s.client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object from S3 (bucket: %s, object: %s): %w", bucketName, objectName, err)
	}
	defer obj.Close()

	data, err := io.ReadAll(obj)
	if err != nil {
		return nil, fmt.Errorf("failed to read object data: %w", err)
	}

	return data, nil
}

// GetObjectStream возвращает поток для чтения объекта из S3
// Это более эффективно для больших файлов, так как не загружает весь файл в память
func (s *s3Client) GetObjectStream(ctx context.Context, bucketName, objectName string) (io.ReadCloser, error) {
	if bucketName == "" {
		bucketName = s.bucketName
	}

	// Убираем ведущий слэш из objectName, если он есть
	objectName = strings.TrimPrefix(objectName, "/")

	obj, err := s.client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object from S3 (bucket: %s, object: %s): %w", bucketName, objectName, err)
	}

	return obj, nil
}
