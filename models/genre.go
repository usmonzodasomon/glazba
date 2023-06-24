package models

type Genre struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null; unique"`
	IsActive bool   `json:"-" gorm:"default:true"`
	// Musics   []Music `json:"musics"`
}

type GenreUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}
