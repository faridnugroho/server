package musicdto

import "dumbsound/models"

type MusicResponse struct {
	ID       int           `json:"id"`
	ArtistID int           `json:"artistid" form:"artistid"`
	Artist   models.Artist `json:"artist"`
	Title    string        `json:"title" form:"title" gorm:"type: varchar(255)"`
	Year     string        `json:"year" form:"year" gorm:"type: varchar(255)"`
	Image    string        `json:"image" form:"image" gorm:"type: varchar(255)"`
	Attache  string        `json:"attache" form:"attache" gorm:"type: varchar(255)"`
}
