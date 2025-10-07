package services

import (
	"context"
	"educabot.com/bookshop/externals"
	"educabot.com/bookshop/models"
	"log/slog"
	"sync"
)

var (
	once     sync.Once
	instance *BooksService
)

type BooksService struct {
	external *externals.ExternalServices
}

func NewBooksService(external *externals.ExternalServices) *BooksService {
	once.Do(func() {
		instance = &BooksService{external: external}
	})
	return instance
}

func (bs *BooksService) GetBooks(ctx context.Context) []models.Book {
	books, err := bs.external.GetBooks(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "BooksService: GetBooks failed", "err", err)
		return []models.Book{}
	}
	return books
}
