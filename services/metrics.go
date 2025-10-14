package services

import (
	"context"
	"errors"
	"slices"

	"educabot.com/bookshop/models"
	"educabot.com/bookshop/providers"
)

type MetricService struct {
	booksService providers.BooksProvider
}

func NewMetricService(booksService providers.BooksProvider) *MetricService {
	return &MetricService{booksService: booksService}
}

func (ms *MetricService) GetMetrics(ctx context.Context, query models.GetMetricsRequest) (models.MetricsResponse, error) {
	books, err := ms.booksService.GetBooks(ctx)
	if err != nil {
		return models.MetricsResponse{}, err
	}

	if len(books) == 0 {
		return models.MetricsResponse{}, errors.New("no books available")
	}

	metrics := models.MetricsResponse{
		MeanUnitsSold:        meanUnitsSold(books),
		CheapestBook:         cheapestBook(books).Name,
		BooksWrittenByAuthor: booksWrittenByAuthor(books, query.Author),
	}

	return metrics, nil
}

func meanUnitsSold(books []models.Book) uint {
	var sum uint
	for _, book := range books {
		sum += book.UnitsSold
	}
	return sum / uint(len(books))
}

func cheapestBook(books []models.Book) models.Book {
	return slices.MinFunc(books, func(a, b models.Book) int {
		return int(a.Price - b.Price)
	})
}

func booksWrittenByAuthor(books []models.Book, author string) uint {
	var count uint
	for _, book := range books {
		if book.Author == author {
			count++
		}
	}
	return count
}
