package models

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"not null"`
	Artist      string `json:"artist" gorm:"not null"`
	Duration    uint16 `json:"duration" gorm:"not null"`
	ReleaseData uint16 `json:"release_data" gorm:"not null"`
	Filepath    string `json:"filepath" gorm:"not null"`
}

type RequestProduct struct {
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	ReleaseData string `json:"release_data"`
	Category    string `json:"category"`
	Genre       string `json:"genre"`
}
