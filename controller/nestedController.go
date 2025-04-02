package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type NestedController interface {
	FindBookById(c *gin.Context)
	FindAllBooks(c *gin.Context)
	CreateBook(c *gin.Context)
	FindAuthorById(c *gin.Context)
	FindAllAuthors(c *gin.Context)
	CreateAuthor(c *gin.Context)
}

// @Summary Find book by id
// @Description Find book by id
// @Tags Nested
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Router /nested/books/{id} [get]
func (n *nestedController) FindBookById(c *gin.Context) {
	log.Info().Msg("Called find book by id from Nested Controller")
}

// @Summary Find all books
// @Description Find all books
// @Tags Nested
// @Accept json
// @Produce json
// @Router /nested/books [get]
func (n *nestedController) FindAllBooks(c *gin.Context) {
	log.Info().Msg("Called find all books from Nested Controller")
}


// @Summary Create book
// @Description Create book
// @Tags Nested
// @Accept json
// @Produce json
// @Param book body string true "Book object"
// @Router /nested/createBook [post]
func (n *nestedController) CreateBook(c *gin.Context) {
	log.Info().Msg("Called create book from Nested Controller")
}


// @Summary Find author by id
// @Description Find author by id
// @Tags Nested
// @Accept json
// @Produce json
// @Param id path string true "Author ID"
// @Router /nested/authors/{id} [get]
func (n *nestedController) FindAuthorById(c *gin.Context) {
	log.Info().Msg("Called find author by id from Nested Controller")
}

// @Summary Find all authors
// @Description Find all authors
// @Tags Nested
// @Accept json
// @Produce json
// @Router /nested/authors [get]
func (n *nestedController) FindAllAuthors(c *gin.Context) {
	log.Info().Msg("Called find all authors from Nested Controller")
}


// @Summary Create author
// @Description Create author
// @Tags Nested
// @Accept json
// @Produce json
// @Param author body string true "Author object"
// @Router /nested/createAuthor [post]
func (n *nestedController) CreateAuthor(c *gin.Context) {
	log.Info().Msg("Called create author from Nested Controller")
}

type nestedController struct {
}

func NewNestedController() NestedController {
	return &nestedController{}
}
