package repository

import (
	"github.com/usmonzodasomon/glazba/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user *models.User) (uint, error)
	GetUser(login string) (models.User, error)
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

type Repository struct {
	Authorization
	Genre
	Playlist
	Music
	Artist
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Genre:         NewGenreRepository(db),
		Playlist:      NewPlaylistRepository(db),
		Music:         NewMusicRepository(db),
		Artist:        NewArtistRepository(db),
	}
}
