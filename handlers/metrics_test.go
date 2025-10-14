package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"educabot.com/bookshop/spec/mockImpls"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMetricController_Handle_OK(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Usar mock del servicio directamente
	mockMetricService := mockImpls.NewMockMetricService()
	controller := NewMetricController(mockMetricService)

	r := gin.Default()
	r.GET("/metrics", controller.Handle())

	req := httptest.NewRequest(http.MethodGet, "/metrics?author=Alan+Donovan", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var body map[string]interface{}
	err := json.Unmarshal(res.Body.Bytes(), &body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, 11000, int(body["mean_units_sold"].(float64)))
	assert.Equal(t, "The Go Programming Language", body["cheapest_book"])
	assert.Equal(t, 1, int(body["books_written_by_author"].(float64)))
}

func TestMetricController_Handle_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Usar mock del servicio que retorna error
	mockMetricService := &mockImpls.MockErrorMetricService{}
	controller := NewMetricController(mockMetricService)

	r := gin.Default()
	r.GET("/metrics", controller.Handle())

	req := httptest.NewRequest(http.MethodGet, "/metrics?author=Alan+Donovan", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code)

	var body map[string]interface{}
	err := json.Unmarshal(res.Body.Bytes(), &body)
	assert.NoError(t, err)
	assert.Equal(t, "Internal server error", body["error"])
}
