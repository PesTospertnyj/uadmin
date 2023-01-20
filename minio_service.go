package uadmin

import (
	"context"
	"io"
)

type MinioService interface {
	UploadFile(ctx context.Context, filename, contentType string, size int64, file io.Reader) (string, error)
}
