package main

import (
	"fmt"
	"log"

	"os"

	"github.com/HiogoPariz/files-notez/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := api.SetupRouter()
	port := ":" + os.Getenv("PORT")
	fmt.Println("port", port)
	router.Run(port)
}
