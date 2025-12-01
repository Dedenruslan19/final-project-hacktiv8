package repository

import (
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
)

type GCSStorageRepo interface {
	UploadFile(ctx context.Context, file io.Reader, objectName string) (string, error)
}

type gcsStorageRepo struct {
	client     *storage.Client
	bucketName string
}

func NewGCSStorageRepo(client *storage.Client, bucketName string) GCSStorageRepo {
	return &gcsStorageRepo{
		client:     client,
		bucketName: bucketName,
	}
}

func (r *gcsStorageRepo) UploadFile(ctx context.Context, file io.Reader, objectName string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	bucket := r.client.Bucket(r.bucketName)
	obj := bucket.Object(objectName)

	writer := obj.NewWriter(ctx)

	if _, err := io.Copy(writer, file); err != nil {
		return "", err
	}

	if err := writer.Close(); err != nil {
		return "", err
	}

	// URL public (if bucket is public)
	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", r.bucketName, objectName)

	return url, nil
}
