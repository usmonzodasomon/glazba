package models

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"not null"`
	Artist      string `json:"artist" gorm:"not null"`
	Duration    uint16 `json:"duration" gorm:"not null"`
	ReleaseData uint16 `json:"release_data" gorm:"not null"`
	Filepath    string `json:"filepath" gorm:"not null"`
}

type Category struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null; unique"`
	IsActive bool   `json:"-" gorm:"default:true"`
}
