package repository

import (
	"github.com/usmonzodasomon/glazba/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user *models.User) (uint, error)
	GetUser(login string) (models.User, error)
}

type Category interface {
	CreateCategory(category *models.Category) (uint, error)
	ReadCategory(*[]models.Category) error
	ReadCategoryByName(categoryName string) (models.Category, error)
	UpdateCategory(categoryName, name string) error
	DeleteCategory(categoryName string) error
}

type Repository struct {
	Authorization
	Category
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Category:      NewCategoryRepository(db),
	}
}
