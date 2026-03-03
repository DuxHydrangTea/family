package providers

import "io"

type S3Provider struct{}

func NewS3Provider() *S3Provider {
	return &S3Provider{}
}

func (r *S3Provider) Upload(file io.Reader, fileName string) (string, error) {
	return fileName, nil
}

func (r *S3Provider) PresignedUrl(fileName string) (string, error) {
	return fileName, nil
}
