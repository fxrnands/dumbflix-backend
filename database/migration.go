package database

import (
	"be-dumbflix/models"
	"be-dumbflix/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Film{},
		&models.Category{},
		&models.Transaction{},
		&models.Episode{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
