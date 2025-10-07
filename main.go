package main

import (
	"educabot.com/bookshop/externals"
	"educabot.com/bookshop/services"
	"fmt"
	"log"

	"educabot.com/bookshop/handlers"
	"github.com/gin-gonic/gin"
)

const baseAPI = "/api/v1"

func main() {
	router := gin.New()
	router.SetTrustedProxies(nil)

	bookService := services.NewBooksService(externals.NewExternalServices())

	bookController := handlers.NewBookController(bookService)
	metricsController := handlers.NewMetricController(services.NewMetricService(bookService))

	//metricsHandler := handlers.NewGetMetrics(mockImpls.NewMockBooksProvider())
	router.GET(fmt.Sprintf("%s/metrics", baseAPI), metricsController.Handle())

	router.GET(fmt.Sprintf("%s/books", baseAPI), bookController.Handle())

	err := router.Run(":3000")
	if err != nil {
		log.Printf("Error starting server: %s", err.Error())
		return
	}
	fmt.Println("Starting server on :3000")
}
