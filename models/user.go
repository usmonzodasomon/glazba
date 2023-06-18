package models

import "time"

type User struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Firstname string    `json:"firstname" gorm:"not null"`
	Lastname  string    `json:"lastname" gorm:"not null"`
	Username  string    `json:"username" gorm:"not null; unique"`
	Email     string    `json:"email" gorm:"not nulll; unique"`
	Password  string    `json:"password" gorm:"not null"`
	Role      string    `json:"-" gorm:"default:user;check:role IN ('user', 'admin')"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-" gorm:"index"`
}
