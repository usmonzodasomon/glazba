package db

import (
	"github.com/usmonzodasomon/glazba/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.User{}, &models.Category{}, models.Playlist{}); err != nil {
		return err
	}
	return nil
}
