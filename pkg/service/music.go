package service

import (
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"strings"
	"time"

	"github.com/tcolgate/mp3"
	"github.com/usmonzodasomon/glazba/models"
	"github.com/usmonzodasomon/glazba/pkg/repository"
)

type MusicService struct {
	repos *repository.Repository
}

func NewMusicService(repos *repository.Repository) *MusicService {
	return &MusicService{repos}
}

func (s *MusicService) CreateMusic(music *models.Music) (uint, error) {
	duration, err := MusicDuration(music.File)
	if err != nil {
		return 0, err
	}
	music.Duration = duration
	filePath := fmt.Sprintf("./files/genre_%v/%s_%d.mp3", music.GenreID, music.Title, time.Now().Unix())
	filePath = strings.ReplaceAll(filePath, " ", "_")
	music.Filepath = filePath
	return s.repos.CreateMusic(music)
}

// return the duration of music
func MusicDuration(music *multipart.FileHeader) (uint16, error) {
	file, err := music.Open()
	if err != nil {
		return 0, err
	}
	defer file.Close()
	d := mp3.NewDecoder(file)
	var f mp3.Frame
	skipped := 0

	t := 0.0
	for {

		if err := d.Decode(&f, &skipped); err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}

		t = t + f.Duration().Seconds()
	}
	return uint16(math.Ceil(t)), nil
}
