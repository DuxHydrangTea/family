package services

import (
	"family/internal/models"
	"family/internal/repositories"
)

type UserService struct {
	r *repositories.UserRepository
}

func NewUserService(r *repositories.UserRepository) *UserService {
	return &UserService{
		r: r,
	}
}

func (s *UserService) GetAll() []models.User {
	return s.r.GetAll()
}
