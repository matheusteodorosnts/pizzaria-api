package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/matheusteodorosnts/pizzaria-api/config"
	"github.com/matheusteodorosnts/pizzaria-api/handler"
	"github.com/matheusteodorosnts/pizzaria-api/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitializeConfig()
	handler.InitializeHandler()
	router.Initialize()
}
