package services

import (
	"context"
	"sync"

	"educabot.com/bookshop/models"
	"educabot.com/bookshop/providers"
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

func (bs *BooksService) GetBooks(ctx context.Context) ([]models.Book, error) {
	books, err := bs.provider.GetBooks(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}
