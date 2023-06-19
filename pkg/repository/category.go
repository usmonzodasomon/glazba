package repository

import (
	"github.com/usmonzodasomon/glazba/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (r *CategoryRepository) CreateCategory(category *models.Category) (uint, error) {
	if err := r.db.Create(&category).Error; err != nil {
		return 0, err
	}
	return category.ID, nil
}

func (r *CategoryRepository) ReadCategory(categories *[]models.Category) error {
	if err := r.db.Where("is_active = ?", true).Find(categories).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) ReadCategoryByName(categoryName string) (models.Category, error) {
	var category models.Category
	if err := r.db.Where("name = ? AND is_active = ?", categoryName, true).Take(&category).Error; err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (r *CategoryRepository) UpdateCategory(categoryName, name string) error {
	if err := r.db.Model(&models.Category{}).Where("name = ?", categoryName).Update("name", name).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) DeleteCategory(categoryName string) error {
	if err := r.db.Model(&models.Category{}).Where("name = ?", categoryName).Update("is_active", false).Error; err != nil {
		return err
	}
	return nil
}
