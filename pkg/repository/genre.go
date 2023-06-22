package repository

import (
	"github.com/usmonzodasomon/glazba/models"
	"gorm.io/gorm"
)

type GenreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) *GenreRepository {
	return &GenreRepository{db}
}

func (r *GenreRepository) CreateGenre(genre *models.Genre) (uint, error) {
	if err := r.db.Create(&genre).Error; err != nil {
		return 0, err
	}
	return genre.ID, nil
}

func (r *GenreRepository) ReadGenre(genres *[]models.Genre) error {
	if err := r.db.Where("is_active = ?", true).Find(genres).Error; err != nil {
		return err
	}
	return nil
}

func (r *GenreRepository) ReadGenreById(genreId uint) (models.Genre, error) {
	var genre models.Genre
	if err := r.db.Where("id = ? AND is_active = ?", genreId, true).Take(&genre).Error; err != nil {
		return models.Genre{}, err
	}
	return genre, nil
}

func (r *GenreRepository) UpdateGenre(genreId uint, name string) error {
	if err := r.db.Model(&models.Genre{}).Where("id = ?", genreId).Update("name", name).Error; err != nil {
		return err
	}
	return nil
}

func (r *GenreRepository) DeleteGenre(genreId uint) error {
	if err := r.db.Model(&models.Genre{}).Where("id = ?", genreId).Update("is_active", false).Error; err != nil {
		return err
	}
	return nil
}
