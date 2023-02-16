package database

import (
	"dumbsound/models"
	"dumbsound/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Artist{},
		&models.Music{},
		&models.Transaction{},
		&models.Role{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
