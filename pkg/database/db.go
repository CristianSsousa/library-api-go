package database

import (
	"library-api-go/internal/domain/authors"

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
	DB.AutoMigrate(&authors.Authors{})
	//DB.AutoMigrate(&Books.Books{})
}
