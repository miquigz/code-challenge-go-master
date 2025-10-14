package handlers

import (
	"net/http"

	"educabot.com/bookshop/providers"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService providers.BooksProvider
}

func NewBookController(bookService providers.BooksProvider) *BookController {
	return &BookController{bookService: bookService}
}

func (bc *BookController) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		books, err := bc.bookService.GetBooks(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		ctx.JSON(http.StatusOK, books)
	}
}
