package services

import (
	"context"
	"testing"

	"educabot.com/bookshop/models"
	"educabot.com/bookshop/spec/mockImpls"
	"github.com/stretchr/testify/assert"
)

func TestMetricService_GetMetrics(t *testing.T) {
	mockProvider := mockImpls.NewMockBooksProvider()
	metricService := NewMetricService(mockProvider)

	ctx := context.Background()
	query := models.GetMetricsRequest{
		Author: "Alan Donovan",
	}

	metrics, err := metricService.GetMetrics(ctx, query)
	assert.NoError(t, err)

	// Validaciones
	assert.Equal(t, uint((5000+15000+13000)/3), metrics.MeanUnitsSold)
	assert.Equal(t, "The Go Programming Language", metrics.CheapestBook)
	assert.Equal(t, uint(1), metrics.BooksWrittenByAuthor)
}

func TestMeanUnitsSold(t *testing.T) {
	books := []models.Book{
		{UnitsSold: 5000},
		{UnitsSold: 15000},
		{UnitsSold: 13000},
	}

	mean := meanUnitsSold(books)
	expected := uint((5000 + 15000 + 13000) / 3)

	assert.Equal(t, expected, mean)
}

func TestCheapestBook(t *testing.T) {
	books := []models.Book{
		{Name: "Book A", Price: 50},
		{Name: "Book B", Price: 40},
		{Name: "Book C", Price: 45},
	}

	cheapest := cheapestBook(books)
	assert.Equal(t, "Book B", cheapest.Name)
	assert.Equal(t, uint(40), cheapest.Price)
}

func TestBooksWrittenByAuthor(t *testing.T) {
	books := []models.Book{
		{Author: "Alan Donovan"},
		{Author: "Robert Martin"},
		{Author: "Alan Donovan"},
	}

	count := booksWrittenByAuthor(books, "Alan Donovan")
	assert.Equal(t, uint(2), count)

	count = booksWrittenByAuthor(books, "Robert Martin")
	assert.Equal(t, uint(1), count)

	count = booksWrittenByAuthor(books, "Someone Else")
	assert.Equal(t, uint(0), count)
}

func TestMetricService_GetMetrics_NoBooks(t *testing.T) {
	mockProvider := &mockImpls.MockEmptyBooksProvider{}
	metricService := NewMetricService(mockProvider)

	ctx := context.Background()
	query := models.GetMetricsRequest{
		Author: "Alan Donovan",
	}

	_, err := metricService.GetMetrics(ctx, query)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no books available")
}
