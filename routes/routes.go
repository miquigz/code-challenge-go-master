package routes

import (
	"educabot.com/bookshop/externals"
	"educabot.com/bookshop/handlers"
	"educabot.com/bookshop/services"
	"github.com/gin-gonic/gin"
)

const baseAPI = "/api/v1"

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)

	// Inicialización de dependencias
	bookProvider := externals.NewExternalServices()
	metricService := services.NewMetricService(bookProvider)

	// Controladores
	bookController := handlers.NewBookController(bookProvider)
	metricsController := handlers.NewMetricController(metricService)

	// Rutas API
	api := router.Group(baseAPI)
	{
		api.GET("/books", bookController.Handle())
		api.GET("/metrics", metricsController.Handle())
	}

	return router
}
