package models

type Artist struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null; unique"`
	IsActive bool   `json:"-" gorm:"default:true"`
	// Musics   []Music `json:"musics"`
}

type ArtistUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}
