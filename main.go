package main

import (
	"blog/internal/handlers"
	"blog/internal/routers"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	handlers.InitDatabase()
	routers.InitRouter()
}
