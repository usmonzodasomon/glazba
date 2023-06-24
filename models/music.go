package models

import (
	"mime/multipart"
	"time"
)

type Music struct {
	ID          uint                  `json:"id" gorm:"primaryKey"`
	Title       string                `json:"title" binding:"required" gorm:"not null"`
	ArtistID    uint                  `json:"artist_id" binding:"required" gorm:"not null"`
	GenreID     uint                  `json:"genre_id" binding:"required" gorm:"not null"`
	ReleaseData uint16                `json:"release_data" binding:"required" gorm:"not null"`
	Duration    uint16                `json:"duration" gorm:"not null"`
	Filepath    string                `json:"filepath" gorm:"not null"`
	File        *multipart.FileHeader `form:"file" json:"-" binding:"required" gorm:"-"`
	Artist      Artist                `json:"-" gorm:"foreignKey:ArtistID"`
	Genre       Genre                 `json:"-" gorm:"foreignKey:GenreID"`
	IsActive    bool                  `json:"-" gorm:"default:true"`
	CreatedAt   time.Time             `json:"-"`
	UpdatedAt   time.Time             `json:"-"`
	DeletedAt   time.Time             `json:"-" gorm:"index"`
	// Playlists   []*Playlist           `gorm:"many2many:music_playlist;"`
}

type MusicRequest struct {
	Title       string                `form:"title" bind:"required"`
	ArtistID    uint                  `form:"artist_id" bind:"required"`
	GenreID     uint                  `form:"genre_id" bind:"required"`
	ReleaseData uint16                `form:"release_data" bind:"required"`
	File        *multipart.FileHeader `form:"file" bind:"required"`
}
