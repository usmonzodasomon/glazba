package models

type Playlist struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" binding:"required" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	UserID      uint   `json:"user_id"`
	User        User   `json:"-" gorm:"foreignKey:UserID"`
	// Musics      []*Music `json:"musics" gorm:"many2many:music_playlist;"`
}

type PlaylistUpdateRequest struct {
	Name        *string
	Description *string
}
