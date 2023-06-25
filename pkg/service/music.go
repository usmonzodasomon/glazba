package service

import (
	"github.com/usmonzodasomon/glazba/models"
	"github.com/usmonzodasomon/glazba/pkg/repository"
)

type MusicService struct {
	repos repository.Music
}

func NewMusicService(repos repository.Music) *MusicService {
	return &MusicService{repos}
}

func (s *MusicService) CreateMusic(input *models.MusicRequest, filePath string) (uint, error) {
	var music models.Music
	music.Title = input.Title
	music.ArtistID = input.ArtistID
	music.GenreID = input.GenreID
	music.ReleaseData = input.ReleaseData
	music.File = input.File

	duration, err := MusicDuration(music.File)
	if err != nil {
		return 0, err
	}

	music.Duration = duration
	music.Filepath = filePath
	return s.repos.CreateMusic(&music)
}

func (s *MusicService) GetMusicById(id uint) (models.Music, error) {
	return s.repos.GetMusicById(id)
}
