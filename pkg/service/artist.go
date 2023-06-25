package service

import (
	"github.com/usmonzodasomon/glazba/models"
	"github.com/usmonzodasomon/glazba/pkg/repository"
)

type ArtistService struct {
	repos repository.Artist
}

func NewArtistService(repos repository.Artist) *ArtistService {
	return &ArtistService{repos}
}

func (s *ArtistService) CreateArtist(artist *models.Artist) (uint, error) {
	return s.repos.CreateArtist(artist)
}

func (s *ArtistService) ReadArtist(artist *[]models.Artist) error {
	return s.repos.ReadArtist(artist)
}

func (s *ArtistService) ReadArtistById(artistId uint) (models.Artist, error) {
	return s.repos.ReadArtistById(artistId)
}

func (s *ArtistService) UpdateArtist(artistId uint, name string) error {
	return s.repos.UpdateArtist(artistId, name)
}

func (s *ArtistService) DeleteArtist(artistId uint) error {
	return s.repos.DeleteArtist(artistId)
}
