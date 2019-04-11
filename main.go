package main

import (
	"github.com/joho/godotenv"
	"github.com/posts-api/database"
	"github.com/posts-api/routes"
	"log"
	"os"
)

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	srv := routes.PaveRoutes()
	srv.Run(os.Getenv("PORT"))

}
