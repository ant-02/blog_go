package main

import (
	"blog/handlers"
	"blog/routers"
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
