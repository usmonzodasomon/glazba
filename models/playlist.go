package models

type Playlist struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	UserID      uint   `json:"-"`
}

type PlaylistRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type PlaylistUpdateRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
