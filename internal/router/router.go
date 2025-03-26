package router

import (
	"library-api-go/docs"
	"library-api-go/internal/domain/authors"
	"library-api-go/internal/middleware"
	"library-api-go/pkg/database"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func LoadRouter() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	authorsRepo := authors.NewAuthorsRepository(database.DB)
	authorsService := authors.NewAuthorsService(authorsRepo)
	authorsHandler := authors.NewAuthorsHandler(authorsService)



	r.Use(middleware.ContentTypeMiddleware)
	authorsGroup := r.Group("/api/v1/authors")
	{
		authorsGroup.GET("", authorsHandler.GetAllAuthors)
		authorsGroup.GET("/:id", authorsHandler.GetAuthorById)
		authorsGroup.POST("", authorsHandler.CreateAuthor)
		authorsGroup.PUT("/:id", authorsHandler.UpdateAuthor)
		authorsGroup.DELETE("/:id", authorsHandler.DeleteAuthor)
	}

	//booksGroup := r.Group("/api/v1/books")
	//{
	//	booksGroup.GET("", controllers.GetBooks)
	//	booksGroup.GET("/:id", controllers.GetBookById)
	//	booksGroup.POST("", controllers.CreateBook)
	//	booksGroup.PUT("/:id", controllers.UpdateBook)
	//	booksGroup.DELETE("/:id", controllers.DeleteBook)
	//}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}