package models

type Genre struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name" gorm:"not null; unique"`
	Musics   []Music `json:"musics"`
	IsActive bool    `json:"-" gorm:"default:true"`
}

type GenreUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}
