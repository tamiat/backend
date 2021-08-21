package handlers

//this file is used to handle all business logic

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tamiat/backend/pkg/domain/content"
	"github.com/tamiat/backend/pkg/service"
	"log"
	"net/http"
	"os"
)

func Start() {
	router := mux.NewRouter()
	dbConnection := getDbConnetion()
	ch := ContentHandlers{service.NewContentService(content.NewContentRepositoryDb(dbConnection))}
	router.HandleFunc("/api/v1/contents", ch.getAllContents).Methods(http.MethodGet)
	//get a content by id
	router.Path("/api/v1/content").Queries("id", "{id}").
		HandlerFunc(ch.getContent).Methods(http.MethodGet)

	//get range of contents
	router.Path("/api/v1/contents").Queries("id", "{id}").
		HandlerFunc(ch.getRangeOfContents).Methods(http.MethodGet)

	//post a content
	router.HandleFunc("/api/v1/content", ch.postContent).Methods(http.MethodPost)

	//
	router.Path("/api/v1/content").Queries("id", "{id}").
		HandlerFunc(ch.deleteContent).Methods(http.MethodDelete)

	router.Path("/api/v1/content").Queries("id", "{id}").
		HandlerFunc(ch.updateContent).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func getDbConnetion() *sql.DB{
	dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		os.Getenv("HOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"),
		os.Getenv("USER"),
		os.Getenv("PASS"))
	db, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to conect to db"))
		panic(err)
	}
	log.Println("connected to db ")

	//test connection
	err = db.Ping()
	if err != nil {
		log.Fatal("cannot ping db")
		panic(err)
	}
	log.Println("pinged db")
	return db
}
