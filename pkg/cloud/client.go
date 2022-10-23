package cloud

import (
	"context"
	"io"
)

type BucketClient interface {
	// Creates a new bucket.
	Create(ctx context.Context, bucket string) error
	// Upload a new object to a bucket and returns its URL to view/download.
	UploadObject(ctx context.Context, bucket, fileName string, body io.Reader) (string, error)
}
