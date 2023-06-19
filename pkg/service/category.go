package service

import (
	"github.com/usmonzodasomon/glazba/models"
	"github.com/usmonzodasomon/glazba/pkg/repository"
)

type CategoryService struct {
	repos *repository.Repository
}

func NewCategoryService(repos *repository.Repository) *CategoryService {
	return &CategoryService{repos}
}

func (s *CategoryService) CreateCategory(category *models.Category) (uint, error) {
	return s.repos.CreateCategory(category)
}

func (s *CategoryService) ReadCategory(categories *[]models.Category) error {
	return s.repos.ReadCategory(categories)
}

func (s *CategoryService) ReadCategoryByName(categoryName string) (models.Category, error) {
	return s.repos.ReadCategoryByName(categoryName)
}

func (s *CategoryService) UpdateCategory(categoryName, name string) error {
	return s.repos.UpdateCategory(categoryName, name)
}

func (s *CategoryService) DeleteCategory(categoryName string) error {
	return s.repos.DeleteCategory(categoryName)
}
