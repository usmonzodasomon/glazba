package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/usmonzodasomon/glazba/models"
	"github.com/usmonzodasomon/glazba/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

var signingKey = "owrhe@Q*(h8hrowojwojoe)"

type tokenClaims struct {
	jwt.StandardClaims
	Id   uint   `json:"id"`
	Role string `json:"role"`
}

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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
		},
		user.ID,
		user.Role,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(tokenString string) (uint, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type of *tokenClaims")
	}
	return claims.Id, claims.Role, nil
}
