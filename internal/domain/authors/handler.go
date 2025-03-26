package authors

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthorsHandler struct {
	authorsService *AuthorsService
}

func NewAuthorsHandler(authorsService *AuthorsService) *AuthorsHandler {
	return &AuthorsHandler{authorsService: authorsService}
}

func (h *AuthorsHandler) GetAllAuthors(c *gin.Context) {
	authors, err := h.authorsService.FindAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(authors) == 0 {
		c.JSON(404, gin.H{"error": "Authors not found"})
		return
	}
	responseAuthors := make([]ResponseAuthorsDTO, len(authors))
	for i, author := range authors {
		responseAuthors[i] = author.ToResponseAuthorsDTO()
	}
	c.JSON(200, responseAuthors)
}

func (h *AuthorsHandler) GetAuthorById(c *gin.Context) {
	id := c.Param("id")
	parserdId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	author, err := h.authorsService.FindById(parserdId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, author.ToResponseAuthorsDTO())
}

func (h *AuthorsHandler) CreateAuthor(c *gin.Context) {
	var author RequestAuthorsDTO
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdAuthor, err := h.authorsService.Create(author.ToAuthors())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, createdAuthor.ToResponseAuthorsDTO())
}

func (h *AuthorsHandler) UpdateAuthor(c *gin.Context) {
	var author RequestAuthorsDTO
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	updatedAuthor, err := h.authorsService.Update(author.ToAuthors())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, updatedAuthor.ToResponseAuthorsDTO())
}

func (h *AuthorsHandler) DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	parserdId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	author, err := h.authorsService.FindById(parserdId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = h.authorsService.Delete(author)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
}