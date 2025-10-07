package handlers

import (
	"educabot.com/bookshop/providers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookController struct {
	bookService providers.BooksProvider
}

func NewBookController(bookService providers.BooksProvider) *BookController {
	return &BookController{bookService: bookService}
}

func (bc *BookController) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		books := bc.bookService.GetBooks(ctx)
		ctx.JSON(http.StatusOK, books)
	}
}
