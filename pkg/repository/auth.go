package repository

import (
	"github.com/usmonzodasomon/glazba/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) CreateUser(user *models.User) (uint, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *AuthRepository) GetUser(login string) (models.User, error) {
	var user models.User
	if err := r.db.Where("username = ? OR email = ?", login, login).Take(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
