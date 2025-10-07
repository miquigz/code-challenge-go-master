package services

import (
	"context"
	"educabot.com/bookshop/models"
	"educabot.com/bookshop/providers"
	"sync"
)

var (
	once     sync.Once
	instance *BooksService
)

type BooksService struct {
	provider providers.BooksProvider
}

func NewBooksService(provider providers.BooksProvider) *BooksService {
	once.Do(func() {
		instance = &BooksService{provider: provider}
	})
	return instance
}

func (bs *BooksService) GetBooks(ctx context.Context) []models.Book {
	books := bs.provider.GetBooks(ctx)
	return books
}
