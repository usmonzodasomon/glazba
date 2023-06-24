package repository

import (
	"github.com/usmonzodasomon/glazba/logger"
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
	logger.GetLogger().Debug(music.ArtistID)
	if err := r.db.Create(&music).Error; err != nil {
		return 0, err
	}
	return music.ID, nil
}

func (r *MusicRepository) GetMusic(musics *[]models.Music) error {
	return r.db.Where("is_active = ", true).Find(musics).Error
}

func (r *MusicRepository) GetMusicById(id uint) (models.Music, error) {
	var music models.Music
	if err := r.db.Where("id = ?", id).Take(&music).Error; err != nil {
		return models.Music{}, err
	}
	return music, nil
}
