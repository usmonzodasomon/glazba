package repository

import (
	"github.com/usmonzodasomon/glazba/models"
	"gorm.io/gorm"
)

type MusicRepository struct {
	db *gorm.DB
}

func NewMusicRepository(db *gorm.DB) *MusicRepository {
	return &MusicRepository{db}
}

func (r *MusicRepository) CreateMusic(music *models.Music) (uint, error) {
	if err := r.db.Create(&music).Error; err != nil {
		return 0, err
	}
	return music.ID, nil
}
