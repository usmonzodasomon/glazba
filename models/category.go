package models

type Category struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null; unique"`
	IsActive bool   `json:"-" gorm:"default:true"`
}
