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
	tx := r.db.Begin()
	if user.Firstname != nil {
		if err := tx.Model(models.User{}).Where("id = ?", userId).Update("firstname", user.Firstname).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if user.Lastname != nil {
		if err := tx.Model(models.User{}).Where("id = ?", userId).Update("lastname", user.Lastname).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if user.Email != nil {
		if err := tx.Model(models.User{}).Where("id = ?", userId).Update("email", user.Email).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (r *UserRepository) ChangeUserPassword(password string, userID uint) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("password", password).Error
}

func (r *UserRepository) GetUserById(id uint) (models.User, error) {
	var user models.User
	if err := r.db.Where("ID = ?", id).Take(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
