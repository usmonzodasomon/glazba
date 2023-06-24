package service

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/usmonzodasomon/glazba/models"
	"github.com/usmonzodasomon/glazba/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repos *repository.Repository
}

func NewUserService(repos *repository.Repository) *UserService {
	return &UserService{repos}
}

func (s *UserService) ReadUser(user *models.User, userId uint) error {
	return s.repos.ReadUser(user, userId)
}

func (s *UserService) UpdateUser(user *models.UserUpdate, userId uint) error {
	if user.Email != nil && !govalidator.IsEmail(*user.Email) {
		return errors.New("email is in incorrect format")
	}
	return s.repos.UpdateUser(user, userId)
}

func (s *UserService) ChangeUserPassword(password *models.ChangeUserPasswordData, userID uint) error {
	var user models.User
	if err := s.ReadUser(&user, userID); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password.OldPassword)); err != nil {
		return err
	}

	NewPasswordHash, err := generatePasswordHash(password.NewPassword)
	if err != nil {
		return err
	}
	return s.repos.ChangeUserPassword(NewPasswordHash, userID)
}
