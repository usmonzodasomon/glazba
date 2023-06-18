package repository

import (
	"github.com/usmonzodasomon/glazba/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user *models.User) (uint, error)
	GetUser(user *models.User) (uint, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}