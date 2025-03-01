package main

import (
	"bytes"
	"encoding/json"
	"library-api-go/controllers"
	"library-api-go/database"
	"library-api-go/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var IDAuthor int
var IDBook int

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r
}

func TestGetAllAuthors(t *testing.T) {
	database.ConectionWithDB()
	CreateAuthorMock()
	defer DeleteAuthorMock()
	r := SetupTestRoutes()
	r.GET("/authors", controllers.GetAuthors)
	req, _ := http.NewRequest("GET", "/authors", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetAuthorById(t *testing.T) {
	database.ConectionWithDB()
	CreateAuthorMock()
	defer DeleteAuthorMock()
	r := SetupTestRoutes()
	r.GET("/authors/:id", controllers.GetAuthorById)
	req, _ := http.NewRequest("GET", "/authors/" + strconv.Itoa(IDAuthor), nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

//func TestCreateAuthor(t *testing.T) {
//	r := SetupTestRoutes()
//	r.POST("/authors", controllers.CreateAuthor)
//	author := models.Author{Name: "John Doe",}
//	valueJson, _ := json.Marshal(author)
//	req, _ := http.NewRequest("POST", "/authors", bytes.NewBuffer(valueJson))
//	resp := httptest.NewRecorder()
//	r.ServeHTTP(resp, req)
//	var authorMock models.Author
//	json.Unmarshal(resp.Body.Bytes(), &authorMock)
//	assert.Equal(t,"John Doe", authorMock.Name)
//	assert.Equal(t, http.StatusOK, resp.Code)
//}

func TestUpdateAuthor(t *testing.T) {
	database.ConectionWithDB()
	CreateAuthorMock()
	r := SetupTestRoutes()
	defer DeleteAuthorMock()
	r.PUT("/authors/:id", controllers.UpdateAuthor)
	author := models.Author{Name: "John Doe 2",}
	valueJson, _ := json.Marshal(author)
	req, _ := http.NewRequest("PUT", "/authors/" +strconv.Itoa(IDAuthor), bytes.NewBuffer(valueJson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "O status code não corresponde ao esperado.")
}

func TestDeleteAuthor(t *testing.T) {
	r := SetupTestRoutes()
	database.ConectionWithDB()
	CreateAuthorMock()
	r.DELETE("/authors/:id", controllers.DeleteAuthor)
	req, _ := http.NewRequest("DELETE", "/authors/" + strconv.Itoa(IDAuthor), nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)	
}

func TestGetAllBooks(t *testing.T) {
	database.ConectionWithDB()
	CreateBookMock()
	defer DeleteBookMock()
	r := SetupTestRoutes()
	r.GET("/books", controllers.GetBooks)
	req, _ := http.NewRequest("GET", "/books", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetBookById(t *testing.T) {
	database.ConectionWithDB()
	CreateBookMock()
	defer DeleteBookMock()
	r := SetupTestRoutes()
	r.GET("/books/:id", controllers.GetBookById)
	req, _ := http.NewRequest("GET", "/books/" + strconv.Itoa(IDBook), nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

//func TestCreateBook(t *testing.T) {
//	r := SetupTestRoutes()
//	r.POST("/books", controllers.CreateBook)
//	book := models.Book{Title: "Book 1", AuthorID: 1,}
//	valueJson, _ := json.Marshal(book)
//	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(valueJson))
//	resp := httptest.NewRecorder()
//	r.ServeHTTP(resp, req)
//	var bookMock models.Book
//	json.Unmarshal(resp.Body.Bytes(), &bookMock)
//	assert.Equal(t,"Book 1", bookMock.Title)
//	assert.Equal(t, http.StatusOK, resp.Code)
//}

func TestUpdateBook(t *testing.T) {
	database.ConectionWithDB()
	CreateBookMock()
	defer DeleteBookMock()
	r := SetupTestRoutes()
	r.PUT("/books/:id", controllers.UpdateBook)
	book := models.Book{Title: "Boooook",}
	valueJson, _ := json.Marshal(book)
	req, _ := http.NewRequest("PUT", "/books/" + strconv.Itoa(IDBook), bytes.NewBuffer(valueJson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeleteBook(t *testing.T) {
	r := SetupTestRoutes()
	database.ConectionWithDB()
	CreateBookMock()
	r.DELETE("/books/:id", controllers.DeleteBook)
	req, _ := http.NewRequest("DELETE", "/books/" + strconv.Itoa(IDBook), nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func CreateAuthorMock() {
	author := models.Author{Name: "John Doe",}
	database.DB.Create(&author)
	IDAuthor = author.ID
}

func DeleteAuthorMock() {
	var author models.Author
	database.DB.Delete(&author, IDAuthor)
}

func CreateBookMock() {
	author := models.Author{Name: "John Doe",}
	database.DB.Create(&author)
	book := models.Book{Title: "Book 1", AuthorID: author.ID,}
	database.DB.Create(&book)
	IDBook = book.ID
}

func DeleteBookMock() {
	var book models.Book
	database.DB.Delete(&book, IDBook)
	var author models.Author
	database.DB.Delete(&author, book.AuthorID)
}