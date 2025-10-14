package providers

import (
	"context"

	"educabot.com/bookshop/models"
)

type BooksProvider interface {
	GetBooks(ctx context.Context) ([]models.Book, error)
}

type MetricService interface {
	GetMetrics(ctx context.Context, query models.GetMetricsRequest) (models.MetricsResponse, error)
}
