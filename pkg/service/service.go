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

type Category interface {
	CreateCategory(category *models.Category) (uint, error)
	ReadCategory(*[]models.Category) error
	ReadCategoryByName(nameCategory string) (models.Category, error)
	UpdateCategory(nameCategory, name string) error
	DeleteCategory(nameCategory string) error
}

type Playlist interface {
	CreatePlaylist(playlist *models.Playlist) (uint, error)
	ReadPlaylists(playlists *[]models.Playlist, userId uint) error
	ReadPlaylistById(playlistId, userId uint) (models.Playlist, error)
	UpdatePlaylist(playlistId, userId uint, playlist *models.PlaylistUpdateRequest) error
	DeletePlaylist(playlistId, userId uint) error
}

type Service struct {
	Authorization
	Category
	Playlist
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthUser(repos),
		Category:      NewCategoryService(repos),
		Playlist:      NewPlaylistService(repos),
	}
}
