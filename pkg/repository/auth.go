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
	return 0, nil
}

func (r *AuthRepository) GetUser(user *models.User) (uint, error) {
	return 0, nil
}
