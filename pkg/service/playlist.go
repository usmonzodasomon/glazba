package service

import (
	"errors"

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
	if !s.repos.IsUnique(playlist.UserID, playlist.Name) {
		return 0, errors.New("dublicate name for playlist")
	}
	return s.repos.CreatePlaylist(playlist)
}

func (s *PlaylistService) ReadPlaylists(playlists *[]models.Playlist, userId uint) error {
	return s.repos.ReadPlaylists(playlists, userId)
}

func (s *PlaylistService) ReadPlaylistById(playlistId, userId uint) (models.Playlist, error) {
	return s.repos.ReadPlaylistById(playlistId, userId)
}

func (s *PlaylistService) UpdatePlaylist(playlistId, userID uint, playlist *models.PlaylistUpdateRequest) error {
	if playlist.Name != nil {
		if !s.repos.IsUnique(userID, *playlist.Name) {
			return errors.New("dublicate name for playlist")
		}
	}
	return s.repos.UpdatePlaylist(playlistId, userID, playlist)
}

func (s *PlaylistService) DeletePlaylist(playlistId, userId uint) error {
	return s.repos.DeletePlaylist(playlistId, userId)
}

func (s *PlaylistService) AddPlaylistMusic(userID, playlistID, musicID uint) error {
	playlist, err := s.repos.ReadPlaylistById(playlistID, userID)
	if err != nil {
		return err
	}

	music, err := s.repos.GetMusicById(musicID)
	if err != nil {
		return err
	}

	return s.repos.AddPlaylistMusic(playlist, music)
}
