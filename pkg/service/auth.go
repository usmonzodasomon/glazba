package service

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
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
	auth     repository.Authorization
	playlist repository.Playlist
}

func NewAuthUser(auth repository.Authorization, playlist repository.Playlist) *AuthService {
	return &AuthService{auth, playlist}
}

func (s *AuthService) CreateUser(RegisterData *models.RegisterData) (uint, error) {

	if !govalidator.IsEmail(RegisterData.Email) {
		return 0, errors.New("email is in incorrect format")
	}

	hashPassword, err := generatePasswordHash(RegisterData.Password)
	if err != nil {
		return 0, err
	}

	var User models.User
	User.Username = RegisterData.Username
	User.Email = RegisterData.Email
	User.Password = hashPassword
	UserID, err := s.auth.CreateUser(&User)
	if err != nil {
		return 0, err
	}

	playlist := models.Playlist{
		Name:        "Favorites",
		Description: "Your favorite musics",
		UserID:      UserID,
	}
	_, err = s.playlist.CreatePlaylist(&playlist)
	return UserID, err
}

func (s *AuthService) GenerateToken(loginData *models.LoginData) (string, error) {
	user, err := s.auth.GetUser(loginData.Login)
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
