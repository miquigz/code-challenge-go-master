package handlers

import (
	"educabot.com/bookshop/repositories/mockImpls"
	"educabot.com/bookshop/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBooks_OK(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// 🧩 Usamos el mock provider para crear el servicio
	mockBooksProvider := mockImpls.NewMockBooksProvider()
	bookService := services.NewBooksService(mockBooksProvider)

	// 🧩 Creamos el controlador
	handler := NewBookController(bookService)

	// 🧩 Configuramos el router
	r := gin.Default()
	r.GET("/", handler.Handle())

	// 🧪 Ejecutamos el request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	// 📦 Parseamos el body
	var resBody []map[string]interface{}
	err := json.Unmarshal(res.Body.Bytes(), &resBody)
	assert.NoError(t, err)

	// ✅ Verificaciones
	assert.Equal(t, http.StatusOK, res.Code)
	assert.NotEmpty(t, resBody)
	assert.Equal(t, "The Go Programming Language", resBody[0]["name"])
	assert.Equal(t, "Alan Donovan", resBody[0]["author"])
}
