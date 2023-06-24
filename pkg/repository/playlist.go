package repository

import (
	"github.com/usmonzodasomon/glazba/logger"
	"github.com/usmonzodasomon/glazba/models"
	"gorm.io/gorm"
)

type PlaylistRepository struct {
	db *gorm.DB
}

func NewPlaylistRepository(db *gorm.DB) *PlaylistRepository {
	return &PlaylistRepository{db}
}

func (r *PlaylistRepository) IsUnique(userID uint, name string) bool {
	var count int64
	r.db.Model(models.Playlist{}).Where("user_id = ? AND name = ?", userID, name).Count(&count)
	return count < 1
}

func (r *PlaylistRepository) CreatePlaylist(playlist *models.Playlist) (uint, error) {
	if err := r.db.Create(&playlist).Error; err != nil {
		return 0, err
	}
	return playlist.ID, nil
}

func (r *PlaylistRepository) ReadPlaylists(playlists *[]models.Playlist, userId uint) error {
	logger.GetLogger().Debug(userId)
	if err := r.db.Where("user_id = ?", userId).Find(&playlists).Error; err != nil {
		return err
	}
	return nil
}

func (r *PlaylistRepository) ReadPlaylistById(playlistId, userId uint) (models.Playlist, error) {
	var playlist models.Playlist
	if err := r.db.Where("id = ? AND user_id = ?", playlistId, userId).Take(&playlist).Error; err != nil {
		return models.Playlist{}, err
	}
	return playlist, nil
}

func (r *PlaylistRepository) UpdatePlaylist(playlistId, userId uint, playlist *models.PlaylistUpdateRequest) error {
	if playlist.Name != nil {
		err := r.db.Model(models.Playlist{}).Where("id = ? AND user_id = ?", playlistId, userId).Update("name", playlist.Name).Error
		if err != nil {
			return err
		}
	}

	if playlist.Description != nil {
		err := r.db.Model(models.Playlist{}).Where("id = ? AND user_id = ?", playlistId, userId).Update("description", playlist.Description).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PlaylistRepository) DeletePlaylist(playlistId, userId uint) error {
	if err := r.db.Where("id = ? AND user_id = ?", playlistId, userId).Delete(&models.Playlist{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *PlaylistRepository) AddPlaylistMusic(playlist models.Playlist, music models.Music) error {
	logger.GetLogger().Debug(playlist)
	logger.GetLogger().Debug(music)
	if err := r.db.Model(&playlist).Association("Musics").Append(&music); err != nil {
		return err
	}
	return nil
}
