package models

import "time"

type Artist struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" gorm:"type: varchar(255)"`
	Old         string    `json:"old" gorm:"type: varchar(255)"`
	Type        string    `json:"type" gorm:"type: varchar(255)"`
	StartCareer string    `json:"startcareer" gorm:"type: varchar(255)"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Artist) TableName() string {
	return "artists"
}
