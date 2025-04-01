package routes

import (
	"github.com/TiagoNora/GoCRUDV2/controller"
	"github.com/gin-gonic/gin"
)

func NestedRoutes(engine *gin.Engine) {
	nestedController := controller.NewNestedController()
	engine.GET("/nested/book/:id", nestedController.FindBookById)
	engine.GET("/nested/books", nestedController.FindAllBooks)
	engine.POST("/nested/createBook", nestedController.CreateBook)	
	engine.GET("/nested/author/:id", nestedController.FindAuthorById)
	engine.GET("/nested/authors", nestedController.FindAllAuthors)
	engine.POST("/nested/createAuthor", nestedController.CreateAuthor)	

}