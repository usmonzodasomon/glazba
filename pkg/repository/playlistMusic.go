package repository

import (
	"github.com/usmonzodasomon/glazba/models"
	"gorm.io/gorm"
)

type PlaylistMusicRepository struct {
	db *gorm.DB
}

func NewPlaylistMusicRepository(db *gorm.DB) *PlaylistMusicRepository {
	return &PlaylistMusicRepository{db}
}

func (r *PlaylistMusicRepository) AddPlaylistMusic(playlist models.Playlist, music models.Music) error {
	if err := r.db.Model(&playlist).Association("Musics").Append(&music); err != nil {
		return err
	}
	return nil
}

func (r *PlaylistMusicRepository) GetFavoritePlaylistID(userID uint) (uint, error) {
	var playlist models.Playlist
	if err := r.db.Where("user_id = ? AND name = ?", userID, "Favorites").Take(&playlist).Error; err != nil {
		return 0, err
	}
	return playlist.ID, nil
}

func (r *PlaylistMusicRepository) DeletePlaylistMusic(playlist models.Playlist, music models.Music) error {
	if err := r.db.Model(&playlist).Association("Musics").Delete(&music); err != nil {
		return err
	}
	return nil
}
