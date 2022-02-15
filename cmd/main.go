package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/tamiat/backend/pkg/handlers"

)

func main(){
	//load env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	//start app
	//handlers.Tmp()
	//handlers.Start()
	handlers.TestStart()
}
