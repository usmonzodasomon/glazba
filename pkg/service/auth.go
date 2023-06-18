package service

import (
	"github.com/usmonzodasomon/glazba/models"
	"github.com/usmonzodasomon/glazba/pkg/repository"
)

type AuthService struct {
	repos *repository.Repository
}

func NewAuthUser(repos *repository.Repository) *AuthService {
	return &AuthService{repos}
}

func (s *AuthService) CreateUser(user *models.User) (uint, error) {
	return 0, nil
}

func (s *AuthService) GenerateToken(user *models.User) (string, error) {
	return "", nil
}

func (s *AuthService) ParseToken(tokenString string) (uint, error) {
	return 0, nil
}
