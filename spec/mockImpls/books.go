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
