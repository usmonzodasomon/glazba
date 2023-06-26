package repository

import (
	"github.com/usmonzodasomon/glazba/models"
	"gorm.io/gorm"
)

type ArtistRepository struct {
	db *gorm.DB
}

func NewArtistRepository(db *gorm.DB) *ArtistRepository {
	return &ArtistRepository{db}
}

func (r *ArtistRepository) CreateArtist(artist *models.Artist) (uint, error) {
	if err := r.db.Create(&artist).Error; err != nil {
		return 0, err
	}
	return artist.ID, nil
}

func (r *ArtistRepository) ReadArtist(artist *[]models.Artist, findParam string) error {
	if err := r.db.Where("is_active = ? AND name ILIKE ?", true, "%"+findParam+"%").Find(artist).Error; err != nil {
		return err
	}
	return nil
}

func (r *ArtistRepository) ReadArtistById(artistId uint) ([]models.Music, error) {
	var artist models.Artist
	if err := r.db.Where("id = ? AND is_active = ?", artistId, true).Preload("Musics").Take(&artist).Error; err != nil {
		return nil, err
	}
	return artist.Musics, nil
}

func (r *ArtistRepository) UpdateArtist(artistId uint, name string) error {
	if err := r.db.Model(&models.Artist{}).Where("id = ?", artistId).Update("name", name).Error; err != nil {
		return err
	}
	return nil
}

func (r *ArtistRepository) DeleteArtist(artistId uint) error {
	if err := r.db.Model(&models.Artist{}).Where("id = ?", artistId).Update("is_active", false).Error; err != nil {
		return err
	}
	return nil
}
