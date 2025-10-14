package mockImpls

import (
	"context"
	"errors"

	"educabot.com/bookshop/models"
)

type MockBooksProvider struct{}

func NewMockBooksProvider() *MockBooksProvider {
	return &MockBooksProvider{}
}

func (m *MockBooksProvider) GetBooks(_ context.Context) ([]models.Book, error) {
	return []models.Book{
		{ID: 1, Name: "The Go Programming Language", Author: "Alan Donovan", UnitsSold: 5000, Price: 40},
		{ID: 2, Name: "Clean Code", Author: "Robert C. Martin", UnitsSold: 15000, Price: 50},
		{ID: 3, Name: "The Pragmatic Programmer", Author: "Andrew Hunt", UnitsSold: 13000, Price: 45},
	}, nil
}

// Mock provider que retorna error para testing
type MockErrorBooksProvider struct{}

func (m *MockErrorBooksProvider) GetBooks(ctx context.Context) ([]models.Book, error) {
	return nil, errors.New("database connection failed")
}

// Mock provider que retorna una lista vacía para testing
type MockEmptyBooksProvider struct{}

func (m *MockEmptyBooksProvider) GetBooks(_ context.Context) ([]models.Book, error) {
	return []models.Book{}, nil
}

// Mock MetricService que retorna datos exitosos
type MockMetricService struct{}

func NewMockMetricService() *MockMetricService {
	return &MockMetricService{}
}

func (m *MockMetricService) GetMetrics(ctx context.Context, query models.GetMetricsRequest) (models.MetricsResponse, error) {
	return models.MetricsResponse{
		MeanUnitsSold:        11000,
		CheapestBook:         "The Go Programming Language",
		BooksWrittenByAuthor: 1,
	}, nil
}

// Mock MetricService que retorna error
type MockErrorMetricService struct{}

func (m *MockErrorMetricService) GetMetrics(ctx context.Context, query models.GetMetricsRequest) (models.MetricsResponse, error) {
	return models.MetricsResponse{}, errors.New("service unavailable")
}
