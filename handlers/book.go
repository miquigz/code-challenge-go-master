package handlers

import (
	"educabot.com/bookshop/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookController struct {
	bookService *services.BooksService
}

func NewBookController(bookService *services.BooksService) *BookController {
	return &BookController{bookService: bookService}
}

func (bc *BookController) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		books := bc.bookService.GetBooks(ctx)
		ctx.JSON(http.StatusOK, books)
	}
}
