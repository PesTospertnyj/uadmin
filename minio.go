package uadmin

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
)

type minioService struct {
	client *minio.Client
	config *MinioConfig
}

func NewMinioService(ctx context.Context, config *MinioConfig) (MinioService, error) {
	client, err := connectMinio(ctx, config)
	if err != nil {
		return nil, err
	}
	return &minioService{client: client, config: config}, nil
}

func (s *minioService) UploadFile(ctx context.Context, filename, contentType string, size int64, file io.Reader) (string, error) {
	_, err := s.client.PutObject(ctx, s.config.bucketName, filename, file, size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}

	if s.config.isHttps {
		return fmt.Sprintf("https://%s/%s/%s", s.config.endpoint, s.config.bucketName, filename), nil
	}

	return fmt.Sprintf("http://%s/%s/%s", s.config.endpoint, s.config.bucketName, filename), nil
}

func connectMinio(ctx context.Context, config *MinioConfig) (*minio.Client, error) {
	client, err := minio.New(config.endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.accessKeyId, config.secretAccessKey, ""),
		Secure: config.useSSl,
	})
	if err != nil {
		return nil, err
	}

	exists, err := client.BucketExists(ctx, config.bucketName)
	if err != nil {
		return nil, err
	}

	if !exists {
		err = client.MakeBucket(ctx, config.bucketName, minio.MakeBucketOptions{Region: config.location})
	}

	if config.policy != "" {
		if err := client.SetBucketPolicy(ctx, config.bucketName, config.policy); err != nil {
			return nil, err
		}
	}

	return client, nil
}
