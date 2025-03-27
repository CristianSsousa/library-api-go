package authors

import (
	"library-api-go/pkg/apperros"
	"library-api-go/pkg/response"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthorsHandler struct {
	authorsService ServiceInterface
}

func NewAuthorsHandler(authorsService ServiceInterface) *AuthorsHandler {
	return &AuthorsHandler{authorsService: authorsService}
}

func (h *AuthorsHandler) GetAllAuthors(c *gin.Context) {
	time.Sleep(5 * time.Second)

	authors, err := h.authorsService.FindAll(c.Request.Context())
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, apperros.ErrInternalServerError)
		return
	}

	if len(authors) == 0 {
		response.Success(c, http.StatusOK, []ResponseAuthorsDTO{})
		return
	}
	responseAuthors := make([]ResponseAuthorsDTO, len(authors))
	for i, author := range authors {
		responseAuthors[i] = author.ToResponseAuthorsDTO()
	}
	response.Success(c, http.StatusOK, responseAuthors)
}

func (h *AuthorsHandler) GetAuthorById(c *gin.Context) {
	id := c.Param("id")
	parserdId, err := strconv.Atoi(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, apperros.ErrBadParamInput)
		return
	}
	author, err := h.authorsService.FindById(c.Request.Context(), parserdId)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, apperros.ErrInternalServerError)
		return
	}
	response.Success(c, http.StatusOK, author.ToResponseAuthorsDTO())
}

func (h *AuthorsHandler) CreateAuthor(c *gin.Context) {
	var author RequestAuthorsDTO
	if err := c.ShouldBindJSON(&author); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, apperros.ErrBadParamInput)
		return
	}

	createdAuthor, err := h.authorsService.Create(c.Request.Context(), author.ToAuthors())
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, apperros.ErrInternalServerError)
		return
	}
	response.Success(c, http.StatusCreated, createdAuthor.ToResponseAuthorsDTO())
}

func (h *AuthorsHandler) UpdateAuthor(c *gin.Context) {
	var author RequestAuthorsDTO
	if err := c.ShouldBindJSON(&author); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, apperros.ErrBadParamInput)
		return
	}
	updatedAuthor, err := h.authorsService.Update(c.Request.Context(), author.ToAuthors())
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, apperros.ErrInternalServerError)
		return
	}
	response.Success(c, http.StatusOK, updatedAuthor.ToResponseAuthorsDTO())
}

func (h *AuthorsHandler) DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	parserdId, err := strconv.Atoi(id)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, apperros.ErrBadParamInput)
		return
	}

	author, err := h.authorsService.FindById(c.Request.Context(), parserdId)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, apperros.ErrInternalServerError)
		return
	}
	err = h.authorsService.Delete(c.Request.Context(), author)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, apperros.ErrInternalServerError)
		return
	}
	response.Success(c, http.StatusNoContent, nil)
}