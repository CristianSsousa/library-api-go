package main

import (
	"library-api-go/internal/router"
	"library-api-go/pkg/database"
)

// @title Library API
// @version 1.0
// @description This is a simple library API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://libraryapi.com
// @contact.email
// @license.name MIT
// @host localhost:8080
// @BasePath /
func main() {
	database.ConectionWithDB()
	router.LoadRouter()
}