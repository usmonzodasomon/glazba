package service

import (
	"errors"

	"github.com/usmonzodasomon/glazba/pkg/repository"
)

type LikeService struct {
	user  repository.User
	music repository.Music
	like  repository.Like
}

func NewLikeService(user repository.User, music repository.Music, like repository.Like) *LikeService {
	return &LikeService{
		user:  user,
		music: music,
		like:  like,
	}
}

func (s *LikeService) AddMusicLike(userID, musicID uint) error {
	user, err := s.user.GetUserById(userID)
	if err != nil {
		return err
	}

	music, err := s.music.GetMusicById(musicID)
	if err != nil {
		return err
	}

	if !s.like.HasUserLike(user, music) {
		return s.like.AddMusicLike(user, music)
	}
	return errors.New("you have already liked this music")
}

func (s *LikeService) DeleteMusicLike(userID, musicID uint) error {
	user, err := s.user.GetUserById(userID)
	if err != nil {
		return err
	}

	music, err := s.music.GetMusicById(musicID)
	if err != nil {
		return err
	}

	if s.like.HasUserLike(user, music) {
		return s.like.DeleteMusicLike(user, music)
	}
	return errors.New("you did'n like this music")
}
