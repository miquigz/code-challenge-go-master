package externals

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"sync"

	"educabot.com/bookshop/models"
)

var (
	once     sync.Once
	instance *ExternalServices
)

const baseURLBooksAPI = "https://6781684b85151f714b0aa5db.mockapi.io/api/v1"

type ExternalServices struct {
	client *http.Client
}

func NewExternalServices() *ExternalServices {
	once.Do(func() {
		instance = &ExternalServices{
			client: &http.Client{},
		}
	})
	return instance
}

func (es *ExternalServices) GetBooks(ctx context.Context) ([]models.Book, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/books", baseURLBooksAPI), nil)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create request", "error", err)
		return nil, err
	}

	resp, err := es.client.Do(req)
	if err != nil {
		slog.ErrorContext(ctx, "failed to perform request", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.ErrorContext(ctx, "received non-200 response", "status", resp.StatusCode)
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var books []models.Book
	if err := json.NewDecoder(resp.Body).Decode(&books); err != nil {
		slog.ErrorContext(ctx, "failed to decode response body", "error", err)
		return nil, err
	}

	return books, nil
}
