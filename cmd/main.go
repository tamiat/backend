package main

import (
	"github.com/joho/godotenv"
	"github.com/tamiat/backend/pkg/handlers"
	"log"
)

func main(){
	//load env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	//start app
	handlers.Start()
}
