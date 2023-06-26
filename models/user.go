package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Username  string    `json:"username" gorm:"not null; unique"`
	Email     string    `json:"email" gorm:"not nulll; unique"`
	Password  string    `json:"-" gorm:"not null"`
	Role      string    `json:"-" gorm:"default:user;check:role IN ('user', 'admin')"`
	Likes     []*Music  `json:"-" gorm:"many2many:likes"`
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
