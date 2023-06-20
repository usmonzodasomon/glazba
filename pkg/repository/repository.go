package repository

import (
	"github.com/usmonzodasomon/glazba/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user *models.User) (uint, error)
	GetUser(login string) (models.User, error)
}

type Category interface {
	CreateCategory(category *models.Category) (uint, error)
	ReadCategory(*[]models.Category) error
	ReadCategoryByName(categoryName string) (models.Category, error)
	UpdateCategory(categoryName, name string) error
	DeleteCategory(categoryName string) error
}

type Playlist interface {
	CreatePlaylist(playlist *models.Playlist) (uint, error)
	ReadPlaylists(playlists *[]models.Playlist, userId uint) error
	ReadPlaylistById(playlistId, userId uint) (models.Playlist, error)
	UpdatePlaylist(playlistId, userId uint, playlist *models.PlaylistUpdateRequest) error
	DeletePlaylist(playlistId, userId uint) error
}

type Repository struct {
	Authorization
	Category
	Playlist
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Category:      NewCategoryRepository(db),
		Playlist:      NewPlaylistRepository(db),
	}
}
