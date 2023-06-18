package service

import (
	"github.com/usmonzodasomon/glazba/models"
	"github.com/usmonzodasomon/glazba/pkg/repository"
)

type Authorization interface {
	CreateUser(user *models.RegisterData) (uint, error)
	GenerateToken(loginData *models.LoginData) (string, error)
	ParseToken(tokenString string) (uint, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthUser(repos),
	}
}
