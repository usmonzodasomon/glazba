package models

import (
	"time"

	"github.com/usmonzodasomon/glazba/db"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Username  string    `json:"username" gorm:"not null; unique"`
	Email     string    `json:"email" gorm:"not nulll; unique"`
	Password  string    `json:"-" gorm:"not null"`
	Role      string    `json:"-" gorm:"default:user;check:role IN ('user', 'admin')"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-" gorm:"index"`
}

type RegisterData struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginData struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdate struct {
	Firstname *string
	Lastname  *string
	Email     *string
}

type ChangeUserPasswordData struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	playlist := Playlist{
		Name:        "Favorites",
		Description: "Your favorites music",
		UserID:      user.ID,
	}
	return db.GetDBConn().Create(&playlist).Error
}
