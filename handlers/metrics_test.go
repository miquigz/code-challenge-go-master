package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"educabot.com/bookshop/repositories/mockImpls"
	"educabot.com/bookshop/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMetricController_Handle_OK(t *testing.T) {
	// 🔹 Setear Gin en modo test
	gin.SetMode(gin.TestMode)

	// 🔹 Crear dependencias: provider → service → controller
	mockProvider := mockImpls.NewMockBooksProvider()
	metricService := services.NewMetricService(mockProvider)
	controller := NewMetricController(metricService)

	// 🔹 Setup del router
	r := gin.Default()
	r.GET("/metrics", controller.Handle())

	// 🔹 Simular request HTTP
	req := httptest.NewRequest(http.MethodGet, "/metrics?author=Alan+Donovan", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	// 🔹 Parsear respuesta
	var body map[string]interface{}
	err := json.Unmarshal(res.Body.Bytes(), &body)
	assert.NoError(t, err)

	// 🔹 Aserciones
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, 11000, int(body["mean_units_sold"].(float64)))
	assert.Equal(t, "The Go Programming Language", body["cheapest_book"])
	assert.Equal(t, 1, int(body["books_written_by_author"].(float64)))
}
