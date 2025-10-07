package main

import (
	"educabot.com/bookshop/routes"
	"fmt"
	"log"
)

const baseAPI = "/api/v1"

func main() {
	router := routes.SetupRouter()

	fmt.Println("Starting server on :3000")
	if err := router.Run(":3000"); err != nil {
		log.Fatalf("Error starting server: %s", err.Error())
	}
}
