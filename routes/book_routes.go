package routes

import (
	"BookStore/controllers"
	"github.com/gin-gonic/gin"
)

type BookRouteController struct {
	bookController controllers.BookController
}

func NewBookRouteController(bookController controllers.BookController) BookRouteController {
	return BookRouteController{bookController}
}

func (rc *BookRouteController) BookRoute(rg *gin.RouterGroup) {
	router := rg.Group("/books")
	router.GET("", rc.bookController.GetBooks)
	router.GET("/:id", rc.bookController.GetBook)
	router.POST("", rc.bookController.CreateBook)
	router.DELETE("/:id", rc.bookController.DeleteBook)
	router.PATCH("/:id", rc.bookController.UpdateBook)
}
