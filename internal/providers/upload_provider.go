package providers

import "io"

type UploadProvider interface {
	Upload(file io.Reader, fileName string) (string, error)
	PresignedUrl(fileName string) (string, error)
}
