package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/usmonzodasomon/glazba/models"
	"github.com/usmonzodasomon/glazba/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

var signingKey = "owrhe@Q*(h8hrowojwojoe)"

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

func (s *AuthService) GenerateToken(loginData *models.LoginData) (string, error) {
	user, err := s.repos.GetUser(loginData.Login)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"iss":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 12).Unix(),
	})

	return token.SignedString([]byte(signingKey))

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
