package service

import (
	"io"
	"math"
	"mime/multipart"

	"github.com/tcolgate/mp3"
	"golang.org/x/crypto/bcrypt"
)

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

func generatePasswordHash(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}
