package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/marwangalal746/backend/src/pkg/config"
	"github.com/marwangalal746/backend/src/pkg/handlers"
	"log"
	"net/http"
	"os"
)

func addSomeContents(app *config.AppConfig) {
	insertStatement := `INSERT INTO contents (title, details)VALUES ('first content', 'blablablablabla')
						,('second content', 'jajajajajaja');`
	var err error
	_, err = app.DB.Exec(insertStatement)
	if err != nil {
		panic(err)
	}
}

// loads values from .env into the system
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	r := mux.NewRouter()
	var app *config.AppConfig
	db := config.ConnectDB()
	app = config.NewConfig(db)
	handlers.SetConfig(app)
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {

		}
	}(app.DB)
	addSomeContents(app)
	r.HandleFunc("/content/{id}", handlers.GetContent).Methods("GET")
	fmt.Printf("Starting server at port %s\n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))

}
