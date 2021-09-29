package main

import (
	"github.com/joho/godotenv"
	"github.com/tamiat/backend/pkg/handlers"
	"log"
)

func main(){
	err := godotenv.Load(".env")
	//var startDate = "2021-09-01"
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	handlers.Start()
}
