package service

import (
	"github.com/usmonzodasomon/glazba/models"
	"github.com/usmonzodasomon/glazba/pkg/repository"
)

type Authorization interface {
	CreateUser(user *models.RegisterData) (uint, error)
	GenerateToken(loginData *models.LoginData) (string, error)
	ParseToken(tokenString string) (uint, string, error)
}

type Genre interface {
	CreateGenre(genre *models.Genre) (uint, error)
	ReadGenre(*[]models.Genre) error
	ReadGenreById(genreId uint) (models.Genre, error)
	UpdateGenre(genreId uint, name string) error
	DeleteGenre(genreId uint) error
}

type Playlist interface {
	CreatePlaylist(playlist *models.Playlist) (uint, error)
	ReadPlaylists(playlists *[]models.Playlist, userId uint) error
	ReadPlaylistById(playlistId, userId uint) (models.Playlist, error)
	UpdatePlaylist(playlistId, userId uint, playlist *models.PlaylistUpdateRequest) error
	DeletePlaylist(playlistId, userId uint) error
}

type Artist interface {
	CreateArtist(artist *models.Artist) (uint, error)
	ReadArtist(*[]models.Artist) error
	ReadArtistById(artistId uint) (models.Artist, error)
	UpdateArtist(artistId uint, name string) error
	DeleteArtist(artistId uint) error
}

type Music interface {
	CreateMusic(music *models.Music) (uint, error)
}

type Service struct {
	Authorization
	Genre
	Playlist
	Music
	Artist
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthUser(repos),
		Genre:         NewGenreService(repos),
		Playlist:      NewPlaylistService(repos),
		Music:         NewMusicService(repos),
		Artist:        NewArtistService(repos),
	}
}
