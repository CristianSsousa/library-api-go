package database

import (
	"library-api-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func ConectionWithDB() {
	stringConexao := "user=root dbname=root password=root host=localhost sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao))
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&models.Author{})
	DB.AutoMigrate(&models.Book{})
}
