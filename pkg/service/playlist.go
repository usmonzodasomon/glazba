package service

import (
	"github.com/usmonzodasomon/glazba/models"
	"github.com/usmonzodasomon/glazba/pkg/repository"
)

type PlaylistService struct {
	repos *repository.Repository
}

func NewPlaylistService(repos *repository.Repository) *PlaylistService {
	return &PlaylistService{repos}
}

func (s *PlaylistService) CreatePlaylist(playlist *models.Playlist, userId uint) (uint, error) {
	playlist.UserID = userId
	return s.repos.CreatePlaylist(playlist)
}

func (s *PlaylistService) ReadPlaylists(playlists *[]models.Playlist, userId uint) error {
	return s.repos.ReadPlaylists(playlists, userId)
}

func (s *PlaylistService) ReadPlaylistById(playlistId, userId uint) (models.Playlist, error) {
	return s.repos.ReadPlaylistById(playlistId, userId)
}

func (s *PlaylistService) UpdatePlaylist(playlistId, userId uint, playlist *models.PlaylistUpdateRequest) error {
	return s.repos.UpdatePlaylist(playlistId, userId, playlist)
}

func (s *PlaylistService) DeletePlaylist(playlistId, userId uint) error {
	return s.repos.DeletePlaylist(playlistId, userId)
}
