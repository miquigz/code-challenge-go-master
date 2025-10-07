package services

import (
	"context"
	"educabot.com/bookshop/repositories/mockImpls"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBooks(t *testing.T) {
	mockProvider := mockImpls.NewMockBooksProvider()

	bookService := NewBooksService(mockProvider)

	ctx := context.Background()
	books := bookService.GetBooks(ctx)

	assert.Len(t, books, 3, "Debe devolver 3 libros")

	assert.Equal(t, uint(1), books[0].ID)
	assert.Equal(t, "The Go Programming Language", books[0].Name)
	assert.Equal(t, "Alan Donovan", books[0].Author)
	assert.Equal(t, uint(5000), books[0].UnitsSold)
	assert.Equal(t, uint(40), books[0].Price)

	assert.Equal(t, uint(2), books[1].ID)
	assert.Equal(t, "Clean Code", books[1].Name)
	assert.Equal(t, "Robert C. Martin", books[1].Author)
	assert.Equal(t, uint(15000), books[1].UnitsSold)
	assert.Equal(t, uint(50), books[1].Price)

	assert.Equal(t, uint(3), books[2].ID)
	assert.Equal(t, "The Pragmatic Programmer", books[2].Name)
	assert.Equal(t, "Andrew Hunt", books[2].Author)
	assert.Equal(t, uint(13000), books[2].UnitsSold)
	assert.Equal(t, uint(45), books[2].Price)
}
