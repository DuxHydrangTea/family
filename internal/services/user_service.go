package services

import (
	"family/internal/models"
	"family/internal/providers"
	"family/internal/repositories"
	"io"
)

type UserService struct {
	r               *repositories.UserRepository
	storageProvider providers.UploadProvider
}

func NewUserService(r *repositories.UserRepository, storageProvider providers.UploadProvider) *UserService {
	return &UserService{
		r:               r,
		storageProvider: storageProvider,
	}
}

func (s *UserService) GetAll() []models.User {
	return s.r.GetAll()
}

func (s *UserService) UploadAvatar(file io.Reader, fileName string) (string, error) {
	return s.storageProvider.Upload(file, fileName)
}
