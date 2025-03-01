package router

import (
	"library-api-go/controllers"
	"library-api-go/middleware"

	docs "library-api-go/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func LoadRouter() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.Use(middleware.ContentTypeMiddleware)
	r.GET("/authors", controllers.GetAuthors)
	r.GET("/authors/:id", controllers.GetAuthorById)
	r.POST("/author", controllers.CreateAuthor)
	r.PUT("/authors/:id", controllers.UpdateAuthor)
	r.DELETE("/authors/:id", controllers.DeleteAuthor)

	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.GetBookById)
	r.POST("/book", controllers.CreateBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}