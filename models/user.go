package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Username  string    `json:"username" gorm:"not null; unique"`
	Email     string    `json:"email" gorm:"not nulll; unique"`
	Password  string    `json:"password" gorm:"not null"`
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
