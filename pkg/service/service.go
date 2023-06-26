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
	Test() ([]models.Music, error)
	CreateGenre(genre *models.Genre) (uint, error)
	ReadGenre(*[]models.Genre) error
	ReadGenreMusicsById(genreId uint) ([]models.Music, error)
	UpdateGenre(genreId uint, name string) error
	DeleteGenre(genreId uint) error
}

type Playlist interface {
	CreatePlaylist(playlist *models.Playlist, userId uint) (uint, error)
	ReadPlaylists(playlists *[]models.Playlist, userId uint, findParam string) error
	ReadPlaylistById(playlistId, userId uint) (models.Playlist, error)
	UpdatePlaylist(playlistId, userId uint, playlist *models.PlaylistUpdateRequest) error
	DeletePlaylist(playlistId, userId uint) error
}

type Artist interface {
	CreateArtist(artist *models.Artist) (uint, error)
	ReadArtist(artists *[]models.Artist, findParam string) error
	ReadArtistById(artistId uint) ([]models.Music, error)
	UpdateArtist(artistId uint, name string) error
	DeleteArtist(artistId uint) error
}

type User interface {
	ReadUser(user *models.User, userId uint) error
	UpdateUser(user *models.UserUpdate, userId uint) error
	ChangeUserPassword(password *models.ChangeUserPasswordData, userID uint) error
}

type Music interface {
	CreateMusic(music *models.MusicRequest, filePath string) (uint, error)
	GetFilepathMusic(id uint) (string, error)
	GetMusicById(id uint) (models.MusicAnswer, error)
	GetMusic(findParam, artistIDParam, genreIDParam, releaseDataMinParam, releaseDataMaxParam, duration string) ([]models.MusicAnswer, error)
	UpdateMusic(id uint, music models.MusicUpdate) error
	DeleteMusic(id uint) error
}

type PlaylistMusic interface {
	GetPlaylistMusics(playlistID uint) ([]*models.Music, error)
	AddPlaylistMusic(userID, playlistID, musicID uint) error
	AddFavoriteMusic(userID, musicID uint) error
	DeletePlaylistMusic(userID, playlistID, musicID uint) error
	DeleteFavoriteMusic(userID, musicID uint) error
}

type Like interface {
	AddMusicLike(userID, musicID uint) error
	DeleteMusicLike(userID, musicID uint) error
}

type Service struct {
	Authorization
	Genre
	Playlist
	Music
	User
	Artist
	PlaylistMusic
	Like
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthUser(repos.Authorization, repos.Playlist),
		Genre:         NewGenreService(repos.Genre),
		Playlist:      NewPlaylistService(repos.Playlist),
		Music:         NewMusicService(repos.Music),
		User:          NewUserService(repos.User),
		Artist:        NewArtistService(repos.Artist),
		PlaylistMusic: NewPlaylistMusicService(repos.Playlist, repos.Music, repos.PlaylistMusic),
		Like:          NewLikeService(repos.User, repos.Music, repos.Like),
	}
}
