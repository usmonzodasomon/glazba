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

func (r *MusicRepository) GetMusic(musics *[]models.Music, findParam string, artistID, genreID, releaseDataMin, releaseDataMax, duration int) error {
	q := r.db.Where("is_active = ? AND title ILIKE ?", true, "%"+findParam+"%")
	q.Where("release_data >= ? AND release_data <= ?", releaseDataMin, releaseDataMax)
	if artistID != 0 {
		q.Where("artist_id = ?", artistID)
	}

	if genreID != 0 {
		q.Where("genre_id = ?", genreID)
	}

	if duration != 0 {
		q.Where("duration <= ?", duration)
	}

	return q.Find(musics).Error
}

func (r *MusicRepository) GetMusicById(id uint) (models.Music, error) {
	var music models.Music
	if err := r.db.Where("id = ? AND is_active = ?", id, true).Take(&music).Error; err != nil {
		return models.Music{}, err
	}
	return music, nil
}

func (r *MusicRepository) UpdateMusic(id uint, music models.MusicUpdate) error {
	tx := r.db.Begin()
	if music.Title != nil {
		if err := tx.Model(models.Music{}).Where("id = ?", id).Update("title", music.Title).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if music.ArtistID != nil {
		if err := tx.Model(models.Music{}).Where("id = ?", id).Update("artist_id", music.ArtistID).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if music.GenreID != nil {
		if err := tx.Model(models.Music{}).Where("id = ?", id).Update("genre_id", music.GenreID).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if music.ReleaseData != nil {
		if err := tx.Model(models.Music{}).Where("id = ?", id).Update("release_data", music.ReleaseData).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *MusicRepository) DeleteMusic(id uint) error {
	if err := r.db.Model(&models.Music{}).Where("id = ?", id).Update("is_active", false).Error; err != nil {
		return err
	}
	return nil
}
