package service

import (
	"github.com/usmonzodasomon/glazba/pkg/repository"
)

type PlaylistMusictService struct {
	playlist      repository.Playlist
	music         repository.Music
	playlistMusic repository.PlaylistMusic
}

func NewPlaylistMusicService(playlist repository.Playlist, music repository.Music, playlistMisic repository.PlaylistMusic) *PlaylistMusictService {
	return &PlaylistMusictService{playlist: playlist, music: music, playlistMusic: playlistMisic}
}

func (s *PlaylistMusictService) AddPlaylistMusic(userID, playlistID, musicID uint) error {
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

func (s *PlaylistMusictService) AddFavoriteMusic(userID, musicID uint) error {
	playlistID, err := s.playlistMusic.GetFavoritePlaylistID(userID)
	if err != nil {
		return err
	}
	return s.AddPlaylistMusic(userID, playlistID, musicID)
}

func (s *PlaylistMusictService) DeletePlaylistMusic(userID, playlistID, musicID uint) error {
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

func (s *PlaylistMusictService) DeleteFavoriteMusic(userID, musicID uint) error {
	playlistID, err := s.playlistMusic.GetFavoritePlaylistID(userID)
	if err != nil {
		return err
	}
	return s.DeletePlaylistMusic(userID, playlistID, musicID)
}
