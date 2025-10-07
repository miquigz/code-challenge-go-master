package services

import (
	"context"
	"educabot.com/bookshop/models"
	"educabot.com/bookshop/providers"
	"slices"
)

type MetricService struct {
	booksService providers.BooksProvider
}

func NewMetricService(booksService providers.BooksProvider) *MetricService {
	return &MetricService{booksService: booksService}
}

func (ms *MetricService) GetMetrics(ctx context.Context, query models.GetMetricsRequest) (metrics models.MetricsResponse) {
	books := ms.booksService.GetBooks(ctx)
	metrics.MeanUnitsSold = meanUnitsSold(ctx, books)
	metrics.CheapestBook = cheapestBook(ctx, books).Name
	metrics.BooksWrittenByAuthor = booksWrittenByAuthor(ctx, books, query.Author)
	return
}

func meanUnitsSold(_ context.Context, books []models.Book) uint {
	var sum uint
	for _, book := range books {
		sum += book.UnitsSold
	}
	return sum / uint(len(books))
}

func cheapestBook(_ context.Context, books []models.Book) models.Book {
	return slices.MinFunc(books, func(a, b models.Book) int {
		return int(a.Price - b.Price)
	})
}

func booksWrittenByAuthor(_ context.Context, books []models.Book, author string) uint {
	var count uint
	for _, book := range books {
		if book.Author == author {
			count++
		}
	}
	return count
}
