package externals

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
	log.Println("URl:", fmt.Sprintf("%s/books", baseURLBooksAPI))
	if err != nil {
		return nil, err
	}

	resp, err := es.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var books []models.Book
	if err := json.NewDecoder(resp.Body).Decode(&books); err != nil {
		return nil, err
	}

	return books, nil
}
