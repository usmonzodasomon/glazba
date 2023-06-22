package models

type Artist struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name" gorm:"not null; unique"`
	Musics   []Music `json:"musics"`
	IsActive bool    `json:"-" gorm:"default:true"`
}

type ArtistUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}
