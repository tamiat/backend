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

//This function is used to add some content to db to test the api
func addSomeContents(app *config.AppConfig) {
	//query of insertion
	insertStatement := `INSERT INTO contents (title, details)VALUES ('first content', 'blablablablabla')
						,('second content', 'jajajajajaja');`
	var err error
	//executing the query
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
	//addSomeContents(app)
	//if id=5 the url will be example.com/api/v0/content?id=5
	r.Host("localhost"+ os.Getenv("PORT")).Path("/api/v0/content").Queries("id","{id}").
		HandlerFunc(handlers.GetContent).Name("GetContent")
	r.HandleFunc("/objects", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello! Parameters: %v", r.URL.Query())
	})
	fmt.Printf("Starting server at port %s\n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))


}
