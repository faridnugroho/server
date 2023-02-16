package artistdto

type ArtistResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name" gorm:"type: varchar(255)"`
	Old         string `json:"old" gorm:"type: varchar(255)"`
	Type        string `json:"type" gorm:"type: varchar(255)"`
	StartCareer string `json:"startCareer" gorm:"type: varchar(255)"`
}
