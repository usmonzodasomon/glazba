package repository

import (
	"github.com/usmonzodasomon/glazba/logger"
	"github.com/usmonzodasomon/glazba/models"
	"gorm.io/gorm"
)

type LikeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) *LikeRepository {
	return &LikeRepository{db}
}

func (r *LikeRepository) AddMusicLike(user models.User, music models.Music) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Model(&user).Association("Likes").Append(&music); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(music).Update("likes_count", gorm.Expr("likes_count + ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *LikeRepository) HasUserLike(user models.User, music models.Music) bool {
	if err := r.db.Preload("Likes").Find(&user).Error; err != nil {
		return false
	}
	for _, v := range user.Likes {
		if v.ID == music.ID {
			logger.GetLogger().Debug(v.ID)
			return true
		}
	}
	return false
}

func (r *LikeRepository) DeleteMusicLike(user models.User, music models.Music) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Model(&user).Association("Likes").Delete(&music); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(music).Update("likes_count", gorm.Expr("likes_count - ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
