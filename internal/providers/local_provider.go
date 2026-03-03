package providers

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

type LocalProvider struct {
	baseDir string
}

func NewLocalProvider() *LocalProvider {
	return &LocalProvider{
		baseDir: "storage",
	}
}

func (l *LocalProvider) Upload(file io.Reader, pathName string) (string, error) {
	fullPath := filepath.Join(l.baseDir, pathName)

	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0777); err != nil {
		return "", err
	}

	out, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}

	relativePath := "/" + fullPath

	return relativePath, nil
}

func (l *LocalProvider) PresignedUrl(fileName string) (string, error) {
	return "", errors.New("LocalProvider does not support PresignedUrl")
}
