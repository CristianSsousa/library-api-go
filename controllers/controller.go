package controllers

import (
	"library-api-go/database"
	"library-api-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all authors
// @Description Get all authors
// @Tags Authors
// @Accept json
// @Produce json
// @Success 200 {object} []models.Author
// @Router /authors [get]
func GetAuthors(c *gin.Context) {
	var authors []models.Author
	database.DB.Find(&authors)
	c.JSON(http.StatusOK, authors)
}

// @Summary Get author by id
// @Description Get author by id
// @Tags Authors
// @Accept json
// @Produce json
// @Param id path int true "Author ID"
// @Success 200 {object} models.Author
// @Failure 404 {object} string
// @Router /authors/{id} [get]
func GetAuthorById(c *gin.Context) {
	var author models.Author
	id := c.Param("id")
	database.DB.First(&author, id)
	if author.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	c.JSON(http.StatusOK, author)
}

// @Summary Create author
// @Description Create author
// @Tags Authors
// @Accept json
// @Produce json
// @Param author body models.Author true "Dados do author"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /authors [post]
func CreateAuthor(c * gin.Context) {
	var author models.Author
	
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&author)
	c.JSON(http.StatusOK, gin.H{"message": "Author created successfully"})

}

// @Summary Update author
// @Description Update author
// @Tags Authors
// @Accept json
// @Produce json
// @Param id path int true "Author ID"
// @Param author body models.Author true "Dados do author"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /authors/{id} [put]
func UpdateAuthor(c * gin.Context) {
	var author models.Author
	id := c.Param("id")
	database.DB.First(&author, id)
	if author.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&author)
	c.JSON(http.StatusOK, gin.H{"message": "Author updated successfully"})
}

// @Summary Delete author
// @Description Delete author
// @Tags Authors
// @Accept json
// @Produce json
// @Param id path int true "Author ID"
// @Success 200 {object} string
// @Failure 404 {object} string
// @Router /authors/{id} [delete]
func DeleteAuthor(c * gin.Context) {
	var author models.Author
	var books []models.Book
	id := c.Param("id")
	database.DB.First(&author, id)
	if author.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	database.DB.Find(&books, "author_id = ?", author.ID)
	if len(books) > 0 {
		database.DB.Delete(&books)
	}
	database.DB.Delete(&author)
	c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
}

// @Summary Get all books
// @Description Get all books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {object} []models.Book
// @Router /books [get]
func GetBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Preload("Author").Find(&books)
	c.JSON(http.StatusOK, books)
}

// @Summary Get book by id
// @Description Get book by id
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} string
// @Router /books/{id} [get]
func GetBookById(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	database.DB.First(&book, id)
	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// @Summary Create book
// @Description Create book
// @Tags Books
// @Accept json
// @Produce json
// @Param book body models.Book true "Dados do Book"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /books [post]
func CreateBook(c * gin.Context) {
	var book models.Book
	
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book created successfully"})

}

// @Summary Update book
// @Description Update book
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Dados do Book"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /books/{id} [put]
func UpdateBook(c * gin.Context) {
	var book models.Book
	id := c.Param("id")
	database.DB.First(&book, id)
	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

// @Summary Delete book
// @Description Delete book
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} string
// @Failure 404 {object} string
// @Router /books/{id} [delete]
func DeleteBook(c * gin.Context) {
	var book models.Book
	id := c.Param("id")
	database.DB.First(&book, id)
	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	database.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}