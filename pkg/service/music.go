package service

import (
	"errors"
	"strconv"

	"github.com/asaskevich/govalidator"
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

func (s *MusicService) GetMusicById(id uint) (models.MusicAnswer, error) {
	music, err := s.repos.GetMusicById(id)
	if err != nil {
		return models.MusicAnswer{}, err
	}
	var MusicAnswer models.MusicAnswer
	MusicAnswer.ID = music.ID
	MusicAnswer.Title = music.Title
	MusicAnswer.ReleaseData = music.ReleaseData
	MusicAnswer.Duration = music.Duration
	MusicAnswer.LikesCount = music.LikesCount
	return MusicAnswer, nil
}

func (s *MusicService) GetFilepathMusic(id uint) (string, error) {
	music, err := s.repos.GetMusicById(id)
	if err != nil {
		return "", err
	}
	return music.Filepath, nil
}

func (s *MusicService) GetMusic(findParam, artistIDParam, genreIDParam, releaseDataMinParam, releaseDataMaxParam, durationParam string) ([]models.MusicAnswer, error) {
	if !govalidator.IsNumeric(releaseDataMinParam) || !govalidator.IsNumeric(releaseDataMaxParam) ||
		!govalidator.IsNumeric(artistIDParam) || !govalidator.IsNumeric(genreIDParam) || !govalidator.IsNumeric(durationParam) {
		return nil, errors.New("bad request params")
	}
	if len(releaseDataMinParam) == 0 {
		releaseDataMinParam = "0"
	}

	if len(releaseDataMaxParam) == 0 {
		releaseDataMaxParam = "9999"
	}

	if len(artistIDParam) == 0 {
		artistIDParam = "0"
	}

	if len(genreIDParam) == 0 {
		genreIDParam = "0"
	}

	if len(durationParam) == 0 {
		durationParam = "0"
	}

	releaseDataMin, err := strconv.Atoi(releaseDataMinParam)
	if err != nil {
		return nil, err
	}

	releaseDataMax, err := strconv.Atoi(releaseDataMaxParam)
	if err != nil {
		return nil, err
	}

	artistID, err := strconv.Atoi(artistIDParam)
	if err != nil {
		return nil, err
	}

	genreID, err := strconv.Atoi(genreIDParam)
	if err != nil {
		return nil, err
	}

	duration, err := strconv.Atoi(durationParam)
	if err != nil {
		return nil, err
	}

	var musics []models.Music
	if err := s.repos.GetMusic(&musics, findParam, artistID, genreID, releaseDataMin, releaseDataMax, duration); err != nil {
		return nil, err
	}
	var musicsAnswer []models.MusicAnswer
	for _, v := range musics {
		var MusicAnswer models.MusicAnswer
		MusicAnswer.ID = v.ID
		MusicAnswer.Title = v.Title
		MusicAnswer.ReleaseData = v.ReleaseData
		MusicAnswer.Duration = v.Duration
		MusicAnswer.LikesCount = v.LikesCount
		musicsAnswer = append(musicsAnswer, MusicAnswer)
	}
	return musicsAnswer, nil
}

func (s *MusicService) UpdateMusic(id uint, music models.MusicUpdate) error {
	return s.repos.UpdateMusic(id, music)
}

func (s *MusicService) DeleteMusic(id uint) error {
	return s.repos.DeleteMusic(id)
}
