package service

import (
	"github.com/usmonzodasomon/glazba/models"
	"github.com/usmonzodasomon/glazba/pkg/repository"
)

type PlaylistMusicService struct {
	playlist      repository.Playlist
	music         repository.Music
	playlistMusic repository.PlaylistMusic
}

func NewPlaylistMusicService(playlist repository.Playlist, music repository.Music, playlistMisic repository.PlaylistMusic) *PlaylistMusicService {
	return &PlaylistMusicService{playlist: playlist, music: music, playlistMusic: playlistMisic}
}

func (s *PlaylistMusicService) AddPlaylistMusic(userID, playlistID, musicID uint) error {
	playlist, err := s.playlist.ReadPlaylistById(playlistID, userID)
	if err != nil {
		return err
	}

	music, err := s.music.GetMusicById(musicID)
	if err != nil {
		return err
	}

	return s.playlistMusic.AddPlaylistMusic(playlist, music)
}

func (s *PlaylistMusicService) AddFavoriteMusic(userID, musicID uint) error {
	playlistID, err := s.playlistMusic.GetFavoritePlaylistID(userID)
	if err != nil {
		return err
	}
	return s.AddPlaylistMusic(userID, playlistID, musicID)
}

func (s *PlaylistMusicService) DeletePlaylistMusic(userID, playlistID, musicID uint) error {
	playlist, err := s.playlist.ReadPlaylistById(playlistID, userID)
	if err != nil {
		return err
	}

	music, err := s.music.GetMusicById(musicID)
	if err != nil {
		return err
	}

	return s.playlistMusic.DeletePlaylistMusic(playlist, music)
}

func (s *PlaylistMusicService) DeleteFavoriteMusic(userID, musicID uint) error {
	playlistID, err := s.playlistMusic.GetFavoritePlaylistID(userID)
	if err != nil {
		return err
	}
	return s.DeletePlaylistMusic(userID, playlistID, musicID)
}

func (s *PlaylistMusicService) GetPlaylistMusics(playlistID uint) ([]*models.Music, error) {
	var playlist models.Playlist
	return s.playlistMusic.GetPlaylistMusics(&playlist, playlistID)
}
