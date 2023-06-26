package service

import (
	"github.com/usmonzodasomon/glazba/models"
	"github.com/usmonzodasomon/glazba/pkg/repository"
)

type GenreService struct {
	repos repository.Genre
}

func NewGenreService(repos repository.Genre) *GenreService {
	return &GenreService{repos}
}

func (s *GenreService) CreateGenre(genre *models.Genre) (uint, error) {
	return s.repos.CreateGenre(genre)
}

func (s *GenreService) ReadGenre(genres *[]models.Genre) error {
	return s.repos.ReadGenre(genres)
}

func (s *GenreService) ReadGenreMusicsById(genreId uint) ([]models.Music, error) {
	return s.repos.ReadGenreMusicsById(genreId)
}

func (s *GenreService) UpdateGenre(genreId uint, name string) error {
	return s.repos.UpdateGenre(genreId, name)
}

func (s *GenreService) DeleteGenre(genreId uint) error {
	return s.repos.DeleteGenre(genreId)
}

func (s *GenreService) Test() ([]models.Music, error) {
	return s.repos.Test()
}
