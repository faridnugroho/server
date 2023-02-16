package artistdto

import (
	"time"
)

type CreateArtistRequest struct {
	Name        string    `json:"name" gorm:"type: varchar(255)"`
	Old         string    `json:"old" gorm:"type: varchar(255)"`
	Type        string    `json:"type" gorm:"type: varchar(255)"`
	StartCareer string    `json:"startCareer" gorm:"type: varchar(255)"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateArtistRequest struct {
	Name        string `json:"name" gorm:"type: varchar(255)"`
	Old         string `json:"old" gorm:"type: varchar(255)"`
	Type        string `json:"type" gorm:"type: varchar(255)"`
	StartCareer string `json:"startCareer" gorm:"type: varchar(255)"`
}
