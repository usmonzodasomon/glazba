package service

import (
	"github.com/usmonzodasomon/glazba/models"
	"github.com/usmonzodasomon/glazba/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repos *repository.Repository
}

func NewAuthUser(repos *repository.Repository) *AuthService {
	return &AuthService{repos}
}

func (s *AuthService) CreateUser(RegisterData *models.RegisterData) (uint, error) {
	hashPassword, err := generatePasswordHash(RegisterData.Password)
	if err != nil {
		return 0, err
	}

	var User models.User
	User.Username = RegisterData.Username
	User.Email = RegisterData.Email
	User.Password = hashPassword
	return s.repos.CreateUser(&User)

}

func (s *AuthService) GenerateToken(user *models.User) (string, error) {
	return "", nil
}

func (s *AuthService) ParseToken(tokenString string) (uint, error) {
	return 0, nil
}

func generatePasswordHash(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}
