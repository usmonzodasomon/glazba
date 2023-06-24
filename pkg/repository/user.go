package repository

import (
	"github.com/usmonzodasomon/glazba/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) ReadUser(user *models.User, userId uint) error {
	return r.db.Where("id = ?", userId).Take(user).Error
}

func (r *UserRepository) UpdateUser(user *models.UserUpdate, userId uint) error {
	if user.Firstname != nil {
		if err := r.db.Model(models.User{}).Where("id = ?", userId).Update("firstname", user.Firstname).Error; err != nil {
			return err
		}
	}

	if user.Lastname != nil {
		if err := r.db.Model(models.User{}).Where("id = ?", userId).Update("lastname", user.Lastname).Error; err != nil {
			return err
		}
	}

	if user.Email != nil {
		if err := r.db.Model(models.User{}).Where("id = ?", userId).Update("email", user.Email).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *UserRepository) ChangeUserPassword(password string, userID uint) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("password", password).Error
}
