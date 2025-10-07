package handlers

//
//import (
//	"context"
//	"log/slog"
//	"net/http"
//	"slices"
//
//	"educabot.com/bookshop/models"
//	"educabot.com/bookshop/providers"
//	"github.com/gin-gonic/gin"
//)
//
//type GetMetricsRequest struct {
//	Author string `form:"author"`
//}
//
//func NewGetMetrics(booksProvider providers.BooksProvider) GetMetrics {
//	return GetMetrics{booksProvider}
//}
//
//type GetMetrics struct {
//	booksProvider providers.BooksProvider
//}
//
//func (h GetMetrics) Handle() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		var query GetMetricsRequest
//		err := ctx.ShouldBindQuery(&query)
//		if err != nil {
//			slog.ErrorContext(ctx, "GetMetrics: failed to bind query", "err", err)
//			return
//		}
//
//		ctx.JSON(http.StatusOK, map[string]interface{}{
//			"mean_units_sold":         meanUnitsSold,
//			"cheapest_book":           cheapestBook,
//			"books_written_by_author": booksWrittenByAuthor,
//		})
//	}
//}
